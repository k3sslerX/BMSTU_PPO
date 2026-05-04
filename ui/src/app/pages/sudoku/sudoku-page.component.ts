import { Component, computed, HostListener, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { RouterLink } from '@angular/router';
import { finalize, Observable } from 'rxjs';
import { Driver, MatrixDrivers, MatrixTeams, SudokuCondition, Team } from '../../core/models/api.models';
import { ApiService } from '../../core/services/api.service';
import { AuthService } from '../../core/services/auth.service';
import { readApiError } from '../../core/utils/api-error';

type Mode = 'drivers' | 'teams';
type Matrix = MatrixDrivers | MatrixTeams;
type Entity = Driver | Team;
type CellStatus = 'empty' | 'correct' | 'wrong';

interface CellState {
  query: string;
  status: CellStatus;
  dropdownOpen: boolean;
  locked: boolean;
  selectedId: string;
  selectedName: string;
  message: string;
}

@Component({
  selector: 'app-sudoku-page',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './sudoku-page.component.html',
  styleUrl: './sudoku-page.component.scss'
})
export class SudokuPageComponent implements OnInit, OnDestroy {
  private readonly api = inject(ApiService);
  protected readonly auth = inject(AuthService);

  protected readonly mode = signal<Mode>('drivers');
  protected readonly busy = signal(false);
  protected readonly directoryBusy = signal(false);
  protected readonly feedback = signal<string | null>(null);
  protected readonly driversMatrix = signal<MatrixDrivers | null>(null);
  protected readonly teamsMatrix = signal<MatrixTeams | null>(null);
  protected readonly drivers = signal<Driver[]>([]);
  protected readonly teams = signal<Team[]>([]);
  protected readonly cellStates = signal<Record<string, CellState>>({});

  protected readonly activeMatrix = computed<Matrix | null>(() =>
    this.mode() === 'drivers' ? this.driversMatrix() : this.teamsMatrix()
  );

  protected readonly columnConditions = computed<SudokuCondition[]>(() => this.activeMatrix()?.condition_specs?.[0] ?? []);
  protected readonly rowConditions = computed<SudokuCondition[]>(() => this.activeMatrix()?.condition_specs?.[1] ?? []);
  protected readonly activeField = computed<Entity[][][]>(() => (this.activeMatrix()?.field ?? []) as Entity[][][]);
  protected readonly activeDirectory = computed<Entity[]>(() =>
    this.mode() === 'drivers' ? this.drivers() : this.teams()
  );
  protected readonly emptyCondition: SudokuCondition = {
    label: 'No condition',
    field: '-',
    op: 'eq',
    value: 0
  };
  private activeDropdown: { rowIndex: number; columnIndex: number; anchor: HTMLElement } | null = null;
  private overlay: HTMLElement | null = null;

  ngOnInit(): void {
    if (this.auth.isLoggedIn()) {
      this.load('drivers');
    }
  }

  ngOnDestroy(): void {
    this.removeOverlay();
  }

  protected load(mode: Mode): void {
    if (!this.auth.isLoggedIn()) {
      return;
    }

    this.mode.set(mode);
    this.feedback.set(null);
    this.cellStates.set({});
    this.closeDropdown();
    this.loadDirectory(mode);
    this.busy.set(true);

    const request$: Observable<Matrix> = mode === 'drivers' ? this.api.getDriverSudoku() : this.api.getTeamSudoku();

    request$
      .pipe(finalize(() => this.busy.set(false)))
      .subscribe({
        next: (matrix: Matrix) => {
          if (mode === 'drivers') {
            this.driversMatrix.set(matrix as MatrixDrivers);
          } else {
            this.teamsMatrix.set(matrix as MatrixTeams);
          }
        },
        error: (error: unknown) => {
          this.feedback.set(readApiError(error));
        }
      });
  }

  protected rowCondition(index: number): SudokuCondition {
    return this.rowConditions()[index] ?? this.emptyCondition;
  }

  protected cellState(rowIndex: number, columnIndex: number): CellState {
    return this.cellStates()[this.cellKey(rowIndex, columnIndex)] ?? emptyCellState();
  }

  protected filteredOptions(rowIndex: number, columnIndex: number): Entity[] {
    const query = this.cellState(rowIndex, columnIndex).query;
    const normalized = normalize(query);
    const usedIds = this.usedEntityIds();
    const availableOptions = this.activeDirectory().filter((entity) => !usedIds.has(entity.id));

    if (!normalized) {
      return availableOptions;
    }

    return availableOptions.filter((entity) => searchableValues(entity).some((value) => normalize(value).includes(normalized)));
  }

  protected entitySubtitle(entity: Entity): string {
    return isDriver(entity) ? `${entity.nationality} • ${entity.birthday}` : entity.country;
  }

  protected inputPlaceholder(): string {
    return this.mode() === 'drivers' ? 'Введите пилота' : 'Введите команду';
  }

  protected openDropdown(rowIndex: number, columnIndex: number, anchor: HTMLElement): void {
    const key = this.cellKey(rowIndex, columnIndex);
    this.activeDropdown = { rowIndex, columnIndex, anchor };
    this.cellStates.update((states) => {
      const next = closeDropdowns(states);
      next[key] = {
        ...emptyCellState(),
        ...next[key],
        dropdownOpen: true
      };
      return next;
    });
    this.renderOverlay();
  }

  protected updateCellInput(rowIndex: number, columnIndex: number, query: string, anchor: HTMLElement): void {
    this.activeDropdown = { rowIndex, columnIndex, anchor };
    const exactMatch = resolveEntity(this.activeDirectory(), query);
    const nextState: Partial<CellState> = {
      query,
      selectedId: exactMatch?.id ?? '',
      selectedName: exactMatch?.name ?? '',
      dropdownOpen: true,
      locked: false,
      status: query.trim() ? 'wrong' : 'empty',
      message: query.trim() ? 'Не подходит для этой ячейки' : ''
    };

    this.patchCell(rowIndex, columnIndex, nextState);

    if (exactMatch) {
      this.applyEntityToCell(rowIndex, columnIndex, exactMatch);
      return;
    }

    this.renderOverlay();
  }

  private applyEntityToCell(rowIndex: number, columnIndex: number, entity: Entity): void {
    const isAlreadyUsed = this.usedEntityIds().has(entity.id);
    const fitsCell = this.activeField()[rowIndex]?.[columnIndex]?.some((candidate) => candidate.id === entity.id) ?? false;
    const isCorrect = fitsCell && !isAlreadyUsed;

    this.patchCell(rowIndex, columnIndex, {
      query: entity.name,
      selectedId: entity.id,
      selectedName: entity.name,
      dropdownOpen: false,
      locked: isCorrect,
      status: isCorrect ? 'correct' : 'wrong',
      message: isAlreadyUsed ? 'Уже используется в этой матрице' : 'Не подходит для этой ячейки'
    });
    this.closeDropdown();
  }

  private loadDirectory(mode: Mode): void {
    if ((mode === 'drivers' && this.drivers().length) || (mode === 'teams' && this.teams().length)) {
      return;
    }

    this.directoryBusy.set(true);
    const request$: Observable<Entity[]> = mode === 'drivers' ? this.api.getDrivers() : this.api.getTeams();

    request$
      .pipe(finalize(() => this.directoryBusy.set(false)))
      .subscribe({
        next: (items: Entity[]) => {
          if (mode === 'drivers') {
            this.drivers.set(items as Driver[]);
          } else {
            this.teams.set(items as Team[]);
          }
          this.renderOverlay();
        },
        error: (error: unknown) => {
          this.feedback.set(readApiError(error));
        }
      });
  }

  private patchCell(rowIndex: number, columnIndex: number, patch: Partial<CellState>): void {
    const key = this.cellKey(rowIndex, columnIndex);
    this.cellStates.update((states) => ({
      ...states,
      [key]: {
        ...emptyCellState(),
        ...states[key],
        ...patch
      }
    }));
  }

  private cellKey(rowIndex: number, columnIndex: number): string {
    return `${this.mode()}:${rowIndex}:${columnIndex}`;
  }

  private usedEntityIds(): ReadonlySet<string> {
    const prefix = `${this.mode()}:`;
    const ids = Object.entries(this.cellStates())
      .filter(([key, state]) => key.startsWith(prefix) && state.locked && !!state.selectedId)
      .map(([, state]) => state.selectedId);

    return new Set(ids);
  }

  @HostListener('document:pointerdown', ['$event'])
  protected handleDocumentPointerDown(event: PointerEvent): void {
    const target = event.target as Node | null;
    if (!target || this.overlay?.contains(target) || this.activeDropdown?.anchor.contains(target)) {
      return;
    }

    this.closeDropdown();
  }

  @HostListener('window:resize')
  @HostListener('window:scroll')
  protected repositionDropdown(): void {
    this.renderOverlay();
  }

  private selectOption(rowIndex: number, columnIndex: number, entity: Entity): void {
    this.applyEntityToCell(rowIndex, columnIndex, entity);
  }

  private closeDropdown(): void {
    this.activeDropdown = null;
    this.cellStates.update(closeDropdowns);
    this.removeOverlay();
  }

  private renderOverlay(): void {
    if (!this.activeDropdown) {
      this.removeOverlay();
      return;
    }

    const { rowIndex, columnIndex, anchor } = this.activeDropdown;
    if (!this.cellState(rowIndex, columnIndex).dropdownOpen) {
      this.removeOverlay();
      return;
    }

    const overlay = this.ensureOverlay();
    this.positionOverlay(overlay, anchor);
    overlay.replaceChildren();

    if (this.directoryBusy()) {
      overlay.appendChild(this.createEmptyItem('Загружаем список...'));
      return;
    }

    const options = this.filteredOptions(rowIndex, columnIndex);
    if (!options.length) {
      overlay.appendChild(this.createEmptyItem('Совпадений нет'));
      return;
    }

    for (const option of options) {
      overlay.appendChild(this.createOptionButton(rowIndex, columnIndex, option));
    }
  }

  private ensureOverlay(): HTMLElement {
    if (this.overlay) {
      return this.overlay;
    }

    const overlay = document.createElement('div');
    overlay.className = 'sudoku-combo-overlay';
    document.body.appendChild(overlay);
    this.overlay = overlay;
    return overlay;
  }

  private positionOverlay(overlay: HTMLElement, anchor: HTMLElement): void {
    const rect = anchor.getBoundingClientRect();
    overlay.style.left = `${rect.left}px`;
    overlay.style.top = `${rect.bottom + 7}px`;
    overlay.style.width = `${rect.width}px`;
  }

  private createOptionButton(rowIndex: number, columnIndex: number, entity: Entity): HTMLButtonElement {
    const button = document.createElement('button');
    button.type = 'button';
    button.className = 'sudoku-combo-option';
    button.addEventListener('click', () => this.selectOption(rowIndex, columnIndex, entity));

    const title = document.createElement('strong');
    title.textContent = entity.name;
    button.appendChild(title);

    const subtitle = document.createElement('span');
    subtitle.textContent = this.entitySubtitle(entity);
    button.appendChild(subtitle);

    return button;
  }

  private createEmptyItem(text: string): HTMLSpanElement {
    const item = document.createElement('span');
    item.className = 'sudoku-combo-empty';
    item.textContent = text;
    return item;
  }

  private removeOverlay(): void {
    this.overlay?.remove();
    this.overlay = null;
  }
}

function emptyCellState(): CellState {
  return {
    query: '',
    status: 'empty',
    dropdownOpen: false,
    locked: false,
    selectedId: '',
    selectedName: '',
    message: ''
  };
}

function closeDropdowns(states: Record<string, CellState>): Record<string, CellState> {
  return Object.fromEntries(Object.entries(states).map(([key, state]) => [key, { ...state, dropdownOpen: false }]));
}

function resolveEntity(entities: Entity[], query: string): Entity | null {
  const normalized = normalize(query);
  if (!normalized) {
    return null;
  }

  return entities.find((entity) => normalize(entity.id) === normalized || normalize(entity.name) === normalized) ?? null;
}

function searchableValues(entity: Entity): string[] {
  return isDriver(entity)
    ? [entity.id, entity.name, entity.nationality, entity.birthday]
    : [entity.id, entity.name, entity.country];
}

function isDriver(entity: Entity): entity is Driver {
  return 'nationality' in entity;
}

function normalize(value: string): string {
  return value.trim().toLowerCase();
}
