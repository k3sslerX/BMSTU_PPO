import { Injectable, inject } from '@angular/core';
import { Driver, FavouritesResponse, Team } from '../models/api.models';
import { AuthService } from './auth.service';

@Injectable({ providedIn: 'root' })
export class FavouriteStoreService {
  private readonly auth = inject(AuthService);
  private readonly prefix = 'racingguru.favourites';

  read(): FavouritesResponse {
    const raw = localStorage.getItem(this.storageKey());
    if (!raw) {
      return emptyFavourites();
    }

    try {
      const parsed = JSON.parse(raw) as Partial<FavouritesResponse>;
      return {
        favourite_drivers: Array.isArray(parsed.favourite_drivers) ? parsed.favourite_drivers : [],
        favourite_teams: Array.isArray(parsed.favourite_teams) ? parsed.favourite_teams : []
      };
    } catch {
      return emptyFavourites();
    }
  }

  replace(favourites: FavouritesResponse): void {
    this.write(favourites);
  }

  toggleDriver(driver: Driver): FavouritesResponse {
    const current = this.read();
    const exists = current.favourite_drivers.some((item) => item.id === driver.id);
    const favourite_drivers = exists
      ? current.favourite_drivers.filter((item) => item.id !== driver.id)
      : [...current.favourite_drivers, driver];

    return this.write({
      favourite_drivers,
      favourite_teams: current.favourite_teams
    });
  }

  toggleTeam(team: Team): FavouritesResponse {
    const current = this.read();
    const exists = current.favourite_teams.some((item) => item.id === team.id);
    const favourite_teams = exists
      ? current.favourite_teams.filter((item) => item.id !== team.id)
      : [...current.favourite_teams, team];

    return this.write({
      favourite_drivers: current.favourite_drivers,
      favourite_teams
    });
  }

  private write(favourites: FavouritesResponse): FavouritesResponse {
    localStorage.setItem(this.storageKey(), JSON.stringify(favourites));
    return favourites;
  }

  private storageKey(): string {
    return `${this.prefix}.${this.auth.userId() ?? 'anonymous'}`;
  }
}

function emptyFavourites(): FavouritesResponse {
  return {
    favourite_drivers: [],
    favourite_teams: []
  };
}
