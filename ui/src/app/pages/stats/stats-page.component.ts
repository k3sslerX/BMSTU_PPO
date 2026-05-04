import { Component, HostListener, computed, OnDestroy, inject, OnInit, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, RouterLink, RouterLinkActive } from '@angular/router';
import { finalize } from 'rxjs';
import { Driver, DriverStats, Stats } from '../../core/models/api.models';
import { ApiService } from '../../core/services/api.service';
import { AuthService } from '../../core/services/auth.service';
import { FavouriteStoreService } from '../../core/services/favourite-store.service';
import { readApiError } from '../../core/utils/api-error';

type MetricKey = keyof Stats;
type MetricDirection = 'higher' | 'lower';
type MetricHighlight = 'none' | 'best' | 'tied';

interface Feedback {
  kind: 'error' | 'info';
  text: string;
}

interface MetricDefinition {
  key: MetricKey;
  label: string;
  direction: MetricDirection;
  ignoreZero?: boolean;
}

const MAX_COMPARISON_ITEMS = 4;

@Component({
  selector: 'app-stats-page',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink, RouterLinkActive],
  templateUrl: './stats-page.component.html',
  styleUrl: './stats-page.component.scss'
})
export class StatsPageComponent implements OnInit, OnDestroy {
  private readonly fb = inject(NonNullableFormBuilder);
  private readonly api = inject(ApiService);
  private readonly auth = inject(AuthService);
  private readonly route = inject(ActivatedRoute);
  private readonly favouriteStore = inject(FavouriteStoreService);
  private feedbackTimeoutId: number | null = null;
  private dropdownAnchor: HTMLElement | null = null;
  private overlay: HTMLElement | null = null;

  protected readonly maxComparisonItems = MAX_COMPARISON_ITEMS;
  protected readonly listBusy = signal(false);
  protected readonly statsBusy = signal(false);
  protected readonly dropdownOpen = signal(false);
  protected readonly drivers = signal<Driver[]>([]);
  protected readonly comparison = signal<DriverStats[]>([]);
  protected readonly favouriteDriverIds = signal<ReadonlySet<string>>(new Set());
  protected readonly favouriteBusyId = signal<string | null>(null);
  protected readonly listFeedback = signal<Feedback | null>(null);
  protected readonly feedback = signal<Feedback | null>(null);

  protected readonly metricDefinitions: MetricDefinition[] = [
    { key: 'total_races', label: 'Старты', direction: 'higher' },
    { key: 'total_wins', label: 'Победы', direction: 'higher' },
    { key: 'total_podiums', label: 'Подиумы', direction: 'higher' },
    { key: 'total_points', label: 'Очки', direction: 'higher' },
    { key: 'total_poles', label: 'Поулы', direction: 'higher' },
    { key: 'best_finish', label: 'Лучший финиш', direction: 'lower', ignoreZero: true },
    { key: 'best_qualifying', label: 'Лучшая квалификация', direction: 'lower', ignoreZero: true },
    { key: 'championships_wins', label: 'Титулы', direction: 'higher' },
    {
      key: 'best_championship_position',
      label: 'Лучшее место в чемпионате',
      direction: 'lower',
      ignoreZero: true
    }
  ];

  protected readonly driverForm = this.fb.group({
    driverQuery: [''],
    driverId: ['', [Validators.required]]
  });

  protected readonly driverQuery = toSignal(this.driverForm.controls.driverQuery.valueChanges, {
    initialValue: ''
  });

  protected readonly filteredDrivers = computed(() => filterDrivers(this.drivers(), this.driverQuery()));

  ngOnInit(): void {
    this.loadDrivers();
    this.loadFavouriteDrivers();
  }

  ngOnDestroy(): void {
    this.clearFeedbackTimer();
    this.removeOverlay();
  }

  protected loadDrivers(): void {
    this.listBusy.set(true);
    this.listFeedback.set(null);

    this.api
      .getDrivers()
      .pipe(finalize(() => this.listBusy.set(false)))
      .subscribe({
        next: (drivers) => {
          this.drivers.set(drivers);
          this.addRouteDriverToComparison();
          this.renderOverlay();
        },
        error: (error) => {
          this.listFeedback.set({ kind: 'error', text: readApiError(error) });
        }
      });
  }

  protected addDriverStats(): void {
    this.resolveDriverInput();

    if (this.driverForm.invalid) {
      this.driverForm.markAllAsTouched();
      this.showFeedback({
        kind: 'error',
        text: 'Выберите пилота из списка или введите точный UUID/имя пилота.'
      });
      return;
    }

    const driverId = this.driverForm.getRawValue().driverId.trim();
    const selectedDriver = this.drivers().find((driver) => driver.id === driverId) ?? null;

    this.requestDriverStats(driverId, selectedDriver, true);
  }

  protected removeDriverStats(driverId: string): void {
    this.comparison.update((items) => items.filter((item) => item.driver.id !== driverId));
    this.showFeedback(null);
  }

  protected openDropdown(anchor: HTMLElement): void {
    this.dropdownAnchor = anchor;
    this.dropdownOpen.set(true);
    this.renderOverlay();
  }

  protected syncDriverInput(query: string, anchor: HTMLElement): void {
    this.dropdownAnchor = anchor;
    this.driverForm.controls.driverId.setValue(resolveDriverId(this.drivers(), query) ?? '');
    this.dropdownOpen.set(true);
    this.renderOverlay();
  }

  protected selectDriver(driver: Driver): void {
    this.driverForm.patchValue({
      driverQuery: driver.name,
      driverId: driver.id
    });
    this.closeDropdown();
    this.showFeedback(null);
  }

  protected selectedDriver(): Driver | null {
    const id = this.driverForm.controls.driverId.value;
    return this.drivers().find((driver) => driver.id === id) ?? null;
  }

  protected isDriverFavourite(driverId: string): boolean {
    return this.favouriteDriverIds().has(driverId);
  }

  protected toggleDriverFavourite(result: DriverStats): void {
    if (!this.auth.isLoggedIn()) {
      this.showFeedback({ kind: 'error', text: 'Войдите, чтобы добавлять пилотов в избранное.' });
      return;
    }

    const driver = result.driver;
    const wasFavourite = this.isDriverFavourite(driver.id);
    this.favouriteBusyId.set(driver.id);

    this.api
      .toggleFavouriteDriver({
        id: driver.id,
        name: driver.name,
        nationality: driver.nationality,
        birthday: driver.birthday
      })
      .pipe(finalize(() => this.favouriteBusyId.set(null)))
      .subscribe({
        next: () => {
          const favourites = this.favouriteStore.toggleDriver(driver);
          this.favouriteDriverIds.set(new Set(favourites.favourite_drivers.map((item) => item.id)));
          this.showFeedback(
            {
              kind: 'info',
              text: wasFavourite ? `${driver.name} удалён из избранного.` : `${driver.name} добавлен в избранное.`
            },
            true
          );
        },
        error: (error) => {
          this.showFeedback({ kind: 'error', text: readApiError(error) });
        }
      });
  }

  protected metricValue(stats: Stats, metric: MetricDefinition): number {
    return stats[metric.key];
  }

  protected metricHighlight(stats: Stats, metric: MetricDefinition): MetricHighlight {
    if (this.comparison().length < 2) {
      return 'none';
    }

    const bestValue = this.bestMetricValue(metric);
    if (bestValue === null || this.metricValue(stats, metric) !== bestValue) {
      return 'none';
    }

    return this.bestMetricCount(metric, bestValue) > 1 ? 'tied' : 'best';
  }

  private bestMetricValue(metric: MetricDefinition): number | null {
    const values = this.comparison()
      .map((item) => this.metricValue(item.stats, metric))
      .filter((value) => !metric.ignoreZero || value > 0);

    if (!values.length) {
      return null;
    }

    return metric.direction === 'higher' ? Math.max(...values) : Math.min(...values);
  }

  private bestMetricCount(metric: MetricDefinition, bestValue: number): number {
    return this.comparison()
      .map((item) => this.metricValue(item.stats, metric))
      .filter((value) => (!metric.ignoreZero || value > 0) && value === bestValue).length;
  }

  private loadFavouriteDrivers(): void {
    const cached = this.favouriteStore.read();
    this.favouriteDriverIds.set(new Set(cached.favourite_drivers.map((driver) => driver.id)));

    if (!this.auth.isLoggedIn()) {
      return;
    }

    this.api.getFavourites().subscribe({
      next: ({ favourite_drivers, favourite_teams }) => {
        this.favouriteStore.replace({ favourite_drivers, favourite_teams });
        this.favouriteDriverIds.set(new Set(favourite_drivers.map((driver) => driver.id)));
      },
      error: (error) => {
        if (error?.status === 404) {
          return;
        }

        this.showFeedback({ kind: 'error', text: readApiError(error) });
      }
    });
  }

  private addRouteDriverToComparison(): void {
    const driverId = this.route.snapshot.queryParamMap.get('driver_id')?.trim();
    if (!driverId) {
      return;
    }

    const selectedDriver = this.drivers().find((driver) => driver.id === driverId) ?? null;
    this.requestDriverStats(driverId, selectedDriver, false);
  }

  private requestDriverStats(driverId: string, selectedDriver: Driver | null, clearSelection: boolean): void {
    if (this.comparison().some((item) => item.driver.id === driverId)) {
      this.showFeedback({ kind: 'error', text: 'Этот пилот уже добавлен в сравнение.' });
      return;
    }

    if (this.comparison().length >= MAX_COMPARISON_ITEMS) {
      this.showFeedback({ kind: 'error', text: 'В сравнении может быть максимум 4 пилота.' });
      return;
    }

    this.statsBusy.set(true);
    this.showFeedback(null);

    this.api
      .getDriverStats(driverId)
      .pipe(finalize(() => this.statsBusy.set(false)))
      .subscribe({
        next: ({ driver_stats }) => {
          const hydratedStats = hydrateDriverStats(driver_stats, selectedDriver);
          this.comparison.update((items) => [...items, hydratedStats]);
          this.showFeedback({ kind: 'info', text: `${hydratedStats.driver.name} добавлен в сравнение.` }, true);
          if (clearSelection) {
            this.clearDriverSelection();
          }
        },
        error: (error) => {
          this.showFeedback({ kind: 'error', text: readApiError(error) });
        }
      });
  }

  private resolveDriverInput(): void {
    const query = this.driverForm.getRawValue().driverQuery;
    const currentId = this.driverForm.controls.driverId.value.trim();
    const selectedDriver = this.drivers().find((driver) => driver.id === currentId);

    if (selectedDriver && matchesDriverQuery(selectedDriver, query)) {
      return;
    }

    this.driverForm.controls.driverId.setValue(resolveDriverId(this.drivers(), query) ?? '');
  }

  private clearDriverSelection(): void {
    this.driverForm.patchValue({
      driverQuery: '',
      driverId: ''
    });
    this.closeDropdown();
  }

  @HostListener('document:pointerdown', ['$event'])
  protected handleDocumentPointerDown(event: PointerEvent): void {
    const target = event.target as Node | null;
    if (!target || this.overlay?.contains(target) || this.dropdownAnchor?.contains(target)) {
      return;
    }

    this.closeDropdown();
  }

  @HostListener('window:resize')
  @HostListener('window:scroll')
  protected repositionDropdown(): void {
    this.renderOverlay();
  }

  private closeDropdown(): void {
    this.dropdownOpen.set(false);
    this.dropdownAnchor = null;
    this.removeOverlay();
  }

  private renderOverlay(): void {
    if (!this.dropdownOpen() || !this.dropdownAnchor) {
      this.removeOverlay();
      return;
    }

    const overlay = this.ensureOverlay();
    this.positionOverlay(overlay);
    overlay.replaceChildren();

    if (this.listBusy()) {
      overlay.appendChild(this.createEmptyItem('Загружаем пилотов...'));
      return;
    }

    const query = this.dropdownAnchor instanceof HTMLInputElement ? this.dropdownAnchor.value : this.driverQuery();
    const drivers = filterDrivers(this.drivers(), query);
    if (!drivers.length) {
      overlay.appendChild(this.createEmptyItem('Совпадений нет'));
      return;
    }

    for (const driver of drivers) {
      overlay.appendChild(this.createDriverButton(driver));
    }
  }

  private ensureOverlay(): HTMLElement {
    if (this.overlay) {
      return this.overlay;
    }

    const overlay = document.createElement('div');
    overlay.className = 'stats-combo-overlay';
    document.body.appendChild(overlay);
    this.overlay = overlay;
    return overlay;
  }

  private positionOverlay(overlay: HTMLElement): void {
    if (!this.dropdownAnchor) {
      return;
    }

    const rect = this.dropdownAnchor.getBoundingClientRect();
    overlay.style.left = `${rect.left}px`;
    overlay.style.top = `${rect.bottom + 7}px`;
    overlay.style.width = `${rect.width}px`;
  }

  private createDriverButton(driver: Driver): HTMLButtonElement {
    const button = document.createElement('button');
    button.type = 'button';
    button.className = 'stats-combo-option';
    button.addEventListener('click', () => this.selectDriver(driver));

    const title = document.createElement('strong');
    title.textContent = driver.name;
    button.appendChild(title);

    const subtitle = document.createElement('span');
    subtitle.textContent = `${driver.nationality} • ${driver.birthday}`;
    button.appendChild(subtitle);

    return button;
  }

  private createEmptyItem(text: string): HTMLSpanElement {
    const item = document.createElement('span');
    item.className = 'stats-combo-empty';
    item.textContent = text;
    return item;
  }

  private removeOverlay(): void {
    this.overlay?.remove();
    this.overlay = null;
  }

  private showFeedback(feedback: Feedback | null, autoHide = false): void {
    this.clearFeedbackTimer();
    this.feedback.set(feedback);

    if (feedback && autoHide) {
      this.feedbackTimeoutId = window.setTimeout(() => {
        this.feedback.set(null);
        this.feedbackTimeoutId = null;
      }, 1000);
    }
  }

  private clearFeedbackTimer(): void {
    if (this.feedbackTimeoutId !== null) {
      window.clearTimeout(this.feedbackTimeoutId);
      this.feedbackTimeoutId = null;
    }
  }
}

function filterDrivers(drivers: Driver[], query: string): Driver[] {
  const normalized = normalize(query);
  if (!normalized) {
    return drivers;
  }

  return drivers.filter((driver) =>
    [driver.name, driver.nationality, driver.birthday, driver.id].some((value) => normalize(value).includes(normalized))
  );
}

function resolveDriverId(drivers: Driver[], query: string): string | null {
  const normalized = normalize(query);
  if (!normalized) {
    return null;
  }

  const idMatch = drivers.find((driver) => normalize(driver.id) === normalized);
  if (idMatch) {
    return idMatch.id;
  }

  const nameMatches = drivers.filter((driver) => normalize(driver.name) === normalized);
  return nameMatches.length === 1 ? nameMatches[0].id : null;
}

function matchesDriverQuery(driver: Driver, query: string): boolean {
  const normalized = normalize(query);
  return normalized.length > 0 && (normalize(driver.id) === normalized || normalize(driver.name) === normalized);
}

function hydrateDriverStats(stats: DriverStats, selectedDriver: Driver | null): DriverStats {
  if (!selectedDriver) {
    return stats;
  }

  return {
    ...stats,
    driver: {
      ...selectedDriver,
      ...stats.driver,
      name: stats.driver.name || selectedDriver.name,
      birthday: stats.driver.birthday || selectedDriver.birthday,
      nationality: stats.driver.nationality || selectedDriver.nationality
    }
  };
}

function normalize(value: string): string {
  return value.trim().toLowerCase();
}
