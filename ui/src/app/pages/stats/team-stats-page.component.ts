import { Component, HostListener, computed, OnDestroy, inject, OnInit, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, RouterLink, RouterLinkActive } from '@angular/router';
import { finalize } from 'rxjs';
import { Stats, Team, TeamStats } from '../../core/models/api.models';
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
  selector: 'app-team-stats-page',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink, RouterLinkActive],
  templateUrl: './team-stats-page.component.html',
  styleUrl: './stats-page.component.scss'
})
export class TeamStatsPageComponent implements OnInit, OnDestroy {
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
  protected readonly teams = signal<Team[]>([]);
  protected readonly comparison = signal<TeamStats[]>([]);
  protected readonly favouriteTeamIds = signal<ReadonlySet<string>>(new Set());
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

  protected readonly teamForm = this.fb.group({
    teamQuery: [''],
    teamId: ['', [Validators.required]]
  });

  protected readonly teamQuery = toSignal(this.teamForm.controls.teamQuery.valueChanges, {
    initialValue: ''
  });

  protected readonly filteredTeams = computed(() => filterTeams(this.teams(), this.teamQuery()));

  ngOnInit(): void {
    this.loadTeams();
    this.loadFavouriteTeams();
  }

  ngOnDestroy(): void {
    this.clearFeedbackTimer();
    this.removeOverlay();
  }

  protected loadTeams(): void {
    this.listBusy.set(true);
    this.listFeedback.set(null);

    this.api
      .getTeams()
      .pipe(finalize(() => this.listBusy.set(false)))
      .subscribe({
        next: (teams) => {
          this.teams.set(teams);
          this.addRouteTeamToComparison();
          this.renderOverlay();
        },
        error: (error) => {
          this.listFeedback.set({ kind: 'error', text: readApiError(error) });
        }
      });
  }

  protected addTeamStats(): void {
    this.resolveTeamInput();

    if (this.teamForm.invalid) {
      this.teamForm.markAllAsTouched();
      this.showFeedback({
        kind: 'error',
        text: 'Выберите команду из списка или введите точный UUID/название команды.'
      });
      return;
    }

    const teamId = this.teamForm.getRawValue().teamId.trim();
    const selectedTeam = this.teams().find((team) => team.id === teamId) ?? null;

    this.requestTeamStats(teamId, selectedTeam, true);
  }

  protected removeTeamStats(teamId: string): void {
    this.comparison.update((items) => items.filter((item) => item.team.id !== teamId));
    this.showFeedback(null);
  }

  protected openDropdown(anchor: HTMLElement): void {
    this.dropdownAnchor = anchor;
    this.dropdownOpen.set(true);
    this.renderOverlay();
  }

  protected syncTeamInput(query: string, anchor: HTMLElement): void {
    this.dropdownAnchor = anchor;
    this.teamForm.controls.teamId.setValue(resolveTeamId(this.teams(), query) ?? '');
    this.dropdownOpen.set(true);
    this.renderOverlay();
  }

  protected selectTeam(team: Team): void {
    this.teamForm.patchValue({
      teamQuery: team.name,
      teamId: team.id
    });
    this.closeDropdown();
    this.showFeedback(null);
  }

  protected selectedTeam(): Team | null {
    const id = this.teamForm.controls.teamId.value;
    return this.teams().find((team) => team.id === id) ?? null;
  }

  protected isTeamFavourite(teamId: string): boolean {
    return this.favouriteTeamIds().has(teamId);
  }

  protected toggleTeamFavourite(result: TeamStats): void {
    if (!this.auth.isLoggedIn()) {
      this.showFeedback({ kind: 'error', text: 'Войдите, чтобы добавлять команды в избранное.' });
      return;
    }

    const team = result.team;
    const wasFavourite = this.isTeamFavourite(team.id);
    this.favouriteBusyId.set(team.id);

    this.api
      .toggleFavouriteTeam({
        id: team.id,
        name: team.name,
        country: team.country
      })
      .pipe(finalize(() => this.favouriteBusyId.set(null)))
      .subscribe({
        next: () => {
          const favourites = this.favouriteStore.toggleTeam(team);
          this.favouriteTeamIds.set(new Set(favourites.favourite_teams.map((item) => item.id)));
          this.showFeedback(
            {
              kind: 'info',
              text: wasFavourite ? `${team.name} удалена из избранного.` : `${team.name} добавлена в избранное.`
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

  private loadFavouriteTeams(): void {
    const cached = this.favouriteStore.read();
    this.favouriteTeamIds.set(new Set(cached.favourite_teams.map((team) => team.id)));

    if (!this.auth.isLoggedIn()) {
      return;
    }

    this.api.getFavourites().subscribe({
      next: ({ favourite_drivers, favourite_teams }) => {
        this.favouriteStore.replace({ favourite_drivers, favourite_teams });
        this.favouriteTeamIds.set(new Set(favourite_teams.map((team) => team.id)));
      },
      error: (error) => {
        if (error?.status === 404) {
          return;
        }

        this.showFeedback({ kind: 'error', text: readApiError(error) });
      }
    });
  }

  private addRouteTeamToComparison(): void {
    const teamId = this.route.snapshot.queryParamMap.get('team_id')?.trim();
    if (!teamId) {
      return;
    }

    const selectedTeam = this.teams().find((team) => team.id === teamId) ?? null;
    this.requestTeamStats(teamId, selectedTeam, false);
  }

  private requestTeamStats(teamId: string, selectedTeam: Team | null, clearSelection: boolean): void {
    if (this.comparison().some((item) => item.team.id === teamId)) {
      this.showFeedback({ kind: 'error', text: 'Эта команда уже добавлена в сравнение.' });
      return;
    }

    if (this.comparison().length >= MAX_COMPARISON_ITEMS) {
      this.showFeedback({ kind: 'error', text: 'В сравнении может быть максимум 4 команды.' });
      return;
    }

    this.statsBusy.set(true);
    this.showFeedback(null);

    this.api
      .getTeamStats(teamId)
      .pipe(finalize(() => this.statsBusy.set(false)))
      .subscribe({
        next: ({ team_stats }) => {
          const hydratedStats = hydrateTeamStats(team_stats, selectedTeam);
          this.comparison.update((items) => [...items, hydratedStats]);
          this.showFeedback({ kind: 'info', text: `${hydratedStats.team.name} добавлена в сравнение.` }, true);
          if (clearSelection) {
            this.clearTeamSelection();
          }
        },
        error: (error) => {
          this.showFeedback({ kind: 'error', text: readApiError(error) });
        }
      });
  }

  private resolveTeamInput(): void {
    const query = this.teamForm.getRawValue().teamQuery;
    this.teamForm.controls.teamId.setValue(resolveTeamId(this.teams(), query) ?? '');
  }

  private clearTeamSelection(): void {
    this.teamForm.patchValue({
      teamQuery: '',
      teamId: ''
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
      overlay.appendChild(this.createEmptyItem('Загружаем команды...'));
      return;
    }

    const query = this.dropdownAnchor instanceof HTMLInputElement ? this.dropdownAnchor.value : this.teamQuery();
    const teams = filterTeams(this.teams(), query);
    if (!teams.length) {
      overlay.appendChild(this.createEmptyItem('Совпадений нет'));
      return;
    }

    for (const team of teams) {
      overlay.appendChild(this.createTeamButton(team));
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

  private createTeamButton(team: Team): HTMLButtonElement {
    const button = document.createElement('button');
    button.type = 'button';
    button.className = 'stats-combo-option';
    button.addEventListener('click', () => this.selectTeam(team));

    const title = document.createElement('strong');
    title.textContent = team.name;
    button.appendChild(title);

    const subtitle = document.createElement('span');
    subtitle.textContent = team.country;
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

function filterTeams(teams: Team[], query: string): Team[] {
  const normalized = normalize(query);
  if (!normalized) {
    return teams;
  }

  return teams.filter((team) => [team.name, team.country, team.id].some((value) => normalize(value).includes(normalized)));
}

function resolveTeamId(teams: Team[], query: string): string | null {
  const normalized = normalize(query);
  if (!normalized) {
    return null;
  }

  const match = teams.find((team) => normalize(team.id) === normalized || normalize(team.name) === normalized);
  return match?.id ?? null;
}

function hydrateTeamStats(stats: TeamStats, selectedTeam: Team | null): TeamStats {
  if (!selectedTeam) {
    return stats;
  }

  return {
    ...stats,
    team: {
      ...selectedTeam,
      ...stats.team,
      name: stats.team.name || selectedTeam.name,
      country: stats.team.country || selectedTeam.country
    }
  };
}

function normalize(value: string): string {
  return value.trim().toLowerCase();
}
