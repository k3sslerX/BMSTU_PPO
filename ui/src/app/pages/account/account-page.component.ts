import { Component, computed, inject, OnInit, signal } from '@angular/core';
import { NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { finalize } from 'rxjs';
import { Driver, SudokuCompletionStats, Team } from '../../core/models/api.models';
import { ApiService } from '../../core/services/api.service';
import { AuthService } from '../../core/services/auth.service';
import { FavouriteStoreService } from '../../core/services/favourite-store.service';
import { readApiError } from '../../core/utils/api-error';

interface Feedback {
  kind: 'success' | 'error';
  text: string;
}

const FAVOURITES_PAGE_SIZE = 5;

@Component({
  selector: 'app-account-page',
  standalone: true,
  imports: [ReactiveFormsModule, RouterLink],
  templateUrl: './account-page.component.html',
  styleUrl: './account-page.component.scss'
})
export class AccountPageComponent implements OnInit {
  private readonly fb = inject(NonNullableFormBuilder);
  private readonly api = inject(ApiService);
  private readonly favouriteStore = inject(FavouriteStoreService);

  protected readonly auth = inject(AuthService);
  protected readonly passwordBusy = signal(false);
  protected readonly favouritesBusy = signal(false);
  protected readonly sudokuStatsBusy = signal(false);
  protected readonly favouriteActionBusyId = signal<string | null>(null);
  protected readonly passwordFeedback = signal<Feedback | null>(null);
  protected readonly favouritesFeedback = signal<Feedback | null>(null);
  protected readonly sudokuStatsFeedback = signal<Feedback | null>(null);
  protected readonly sudokuCompletionStats = signal<SudokuCompletionStats>({
    driver_matrices: 0,
    team_matrices: 0
  });
  protected readonly favouriteDrivers = signal<Driver[]>([]);
  protected readonly favouriteTeams = signal<Team[]>([]);
  protected readonly driverPage = signal(0);
  protected readonly teamPage = signal(0);
  protected readonly favouritesPageSize = FAVOURITES_PAGE_SIZE;
  protected readonly driverPageCount = computed(() => pageCount(this.favouriteDrivers().length));
  protected readonly teamPageCount = computed(() => pageCount(this.favouriteTeams().length));
  protected readonly pagedFavouriteDrivers = computed(() => pageItems(this.favouriteDrivers(), this.driverPage()));
  protected readonly pagedFavouriteTeams = computed(() => pageItems(this.favouriteTeams(), this.teamPage()));

  protected readonly passwordForm = this.fb.group({
    password: ['', [Validators.required, Validators.minLength(3)]]
  });

  ngOnInit(): void {
    this.loadProfile();
    this.loadSudokuCompletionStats();
    this.loadFavourites();
  }

  protected changePassword(): void {
    if (this.passwordForm.invalid) {
      this.passwordForm.markAllAsTouched();
      return;
    }

    this.passwordBusy.set(true);
    this.passwordFeedback.set(null);

    const password = this.passwordForm.getRawValue().password;

    this.api
      .changePassword(password)
      .pipe(finalize(() => this.passwordBusy.set(false)))
      .subscribe({
        next: () => {
          this.passwordFeedback.set({
            kind: 'success',
            text: 'Пароль обновлён. JWT остаётся тем же, но новые логины должны использовать новое значение.'
          });
          this.passwordForm.reset();
        },
        error: (error) => {
          this.passwordFeedback.set({
            kind: 'error',
            text: readApiError(error)
          });
        }
      });
  }

  protected loadFavourites(): void {
    this.favouritesBusy.set(true);
    this.favouritesFeedback.set(null);

    const cached = this.favouriteStore.read();
    this.favouriteDrivers.set(cached.favourite_drivers);
    this.favouriteTeams.set(cached.favourite_teams);
    this.clampFavouritePages();

    this.api
      .getFavourites()
      .pipe(finalize(() => this.favouritesBusy.set(false)))
      .subscribe({
        next: ({ favourite_drivers, favourite_teams }) => {
          this.favouriteStore.replace({ favourite_drivers, favourite_teams });
          this.favouriteDrivers.set(favourite_drivers);
          this.favouriteTeams.set(favourite_teams);
          this.clampFavouritePages();
        },
        error: (error) => {
          if (error?.status === 404) {
            return;
          }

          this.favouritesFeedback.set({
            kind: 'error',
            text: readApiError(error)
          });
        }
      });
  }

  protected loadSudokuCompletionStats(): void {
    if (!this.auth.isLoggedIn()) {
      return;
    }

    this.sudokuStatsBusy.set(true);
    this.sudokuStatsFeedback.set(null);

    this.api
      .getSudokuCompletionStats()
      .pipe(finalize(() => this.sudokuStatsBusy.set(false)))
      .subscribe({
        next: (stats) => {
          this.sudokuCompletionStats.set(stats);
        },
        error: (error) => {
          this.sudokuStatsFeedback.set({
            kind: 'error',
            text: readApiError(error)
          });
        }
      });
  }

  protected removeFavouriteDriver(driver: Driver): void {
    this.favouriteActionBusyId.set(`driver:${driver.id}`);
    this.favouritesFeedback.set(null);

    this.api
      .toggleFavouriteDriver({
        id: driver.id,
        name: driver.name,
        nationality: driver.nationality,
        birthday: driver.birthday
      })
      .pipe(finalize(() => this.favouriteActionBusyId.set(null)))
      .subscribe({
        next: () => {
          const favourites = this.favouriteStore.toggleDriver(driver);
          this.favouriteDrivers.set(favourites.favourite_drivers);
          this.favouriteTeams.set(favourites.favourite_teams);
          this.clampFavouritePages();
          this.favouritesFeedback.set({
            kind: 'success',
            text: `${driver.name} удалён из избранного.`
          });
        },
        error: (error) => {
          this.favouritesFeedback.set({
            kind: 'error',
            text: readApiError(error)
          });
        }
      });
  }

  protected removeFavouriteTeam(team: Team): void {
    this.favouriteActionBusyId.set(`team:${team.id}`);
    this.favouritesFeedback.set(null);

    this.api
      .toggleFavouriteTeam({
        id: team.id,
        name: team.name,
        country: team.country
      })
      .pipe(finalize(() => this.favouriteActionBusyId.set(null)))
      .subscribe({
        next: () => {
          const favourites = this.favouriteStore.toggleTeam(team);
          this.favouriteDrivers.set(favourites.favourite_drivers);
          this.favouriteTeams.set(favourites.favourite_teams);
          this.clampFavouritePages();
          this.favouritesFeedback.set({
            kind: 'success',
            text: `${team.name} удалена из избранного.`
          });
        },
        error: (error) => {
          this.favouritesFeedback.set({
            kind: 'error',
            text: readApiError(error)
          });
        }
      });
  }

  protected setDriverPage(page: number): void {
    this.driverPage.set(clampPage(page, this.driverPageCount()));
  }

  protected setTeamPage(page: number): void {
    this.teamPage.set(clampPage(page, this.teamPageCount()));
  }

  protected pageStart(page: number): number {
    return page * FAVOURITES_PAGE_SIZE + 1;
  }

  protected pageEnd(page: number, total: number): number {
    return Math.min(total, (page + 1) * FAVOURITES_PAGE_SIZE);
  }

  private loadProfile(): void {
    if (!this.auth.isLoggedIn()) {
      return;
    }

    this.api.getMe().subscribe({
      next: (profile) => {
        this.auth.setProfile(profile);
      },
      error: () => {
        // Account can still render while the backend is restarting.
      }
    });
  }

  private clampFavouritePages(): void {
    this.setDriverPage(this.driverPage());
    this.setTeamPage(this.teamPage());
  }
}

function pageItems<T>(items: T[], page: number): T[] {
  const start = page * FAVOURITES_PAGE_SIZE;
  return items.slice(start, start + FAVOURITES_PAGE_SIZE);
}

function pageCount(total: number): number {
  return Math.max(1, Math.ceil(total / FAVOURITES_PAGE_SIZE));
}

function clampPage(page: number, totalPages: number): number {
  return Math.max(0, Math.min(page, totalPages - 1));
}
