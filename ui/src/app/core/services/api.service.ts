import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
  AuthUserResponse,
  AdminUser,
  Car,
  CarParticipant,
  Championship,
  Driver,
  DriverStatsResponse,
  FavouritesResponse,
  MatrixDrivers,
  MatrixTeams,
  Race,
  Role,
  Team,
  TeamStatsResponse,
  TokenResponse,
  Track
} from '../models/api.models';
import { getRuntimeConfig } from '../config/runtime-config';

@Injectable({ providedIn: 'root' })
export class ApiService {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = getRuntimeConfig().apiBaseUrl;

  health(): Observable<void> {
    return this.http.get<void>(this.url('/'));
  }

  login(payload: { email: string; password: string }): Observable<TokenResponse> {
    return this.http.post<TokenResponse>(this.url('/login'), payload);
  }

  register(payload: { name: string; email: string; password: string }): Observable<AuthUserResponse> {
    return this.http.post<AuthUserResponse>(this.url('/register'), payload);
  }

  changePassword(password: string): Observable<void> {
    return this.http.post<void>(this.url('/change-password'), { password });
  }

  getMe(): Observable<AuthUserResponse> {
    return this.http.get<AuthUserResponse>(this.url('/me'));
  }

  getDrivers(query = ''): Observable<Driver[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Driver[]>(this.url('/drivers'), { params });
  }

  getTeams(query = ''): Observable<Team[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Team[]>(this.url('/teams'), { params });
  }

  getFavourites(): Observable<FavouritesResponse> {
    return this.http.get<FavouritesResponse>(this.url('/users/favourites'));
  }

  getDriverStats(driverId: string): Observable<DriverStatsResponse> {
    return this.http.get<DriverStatsResponse>(this.url('/stats/driver'), {
      params: new HttpParams().set('driver_id', driverId)
    });
  }

  getTeamStats(teamId: string): Observable<TeamStatsResponse> {
    return this.http.get<TeamStatsResponse>(this.url('/stats/team'), {
      params: new HttpParams().set('team_id', teamId)
    });
  }

  getDriverSudoku(): Observable<MatrixDrivers> {
    return this.http.get<MatrixDrivers>(this.url('/sudoku/drivers'));
  }

  getTeamSudoku(): Observable<MatrixTeams> {
    return this.http.get<MatrixTeams>(this.url('/sudoku/teams'));
  }

  toggleFavouriteDriver(payload: unknown): Observable<void> {
    return this.http.post<void>(this.url('/users/favourite-driver'), payload);
  }

  toggleFavouriteTeam(payload: unknown): Observable<void> {
    return this.http.post<void>(this.url('/users/favourite-team'), payload);
  }

  createDriver(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/drivers'), payload);
  }

  updateDriver(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/drivers'), payload);
  }

  deleteDriver(id: string): Observable<void> {
    return this.http.delete<void>(this.url(`/admin/drivers/${id}`));
  }

  createTeam(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/teams'), payload);
  }

  updateTeam(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/teams'), payload);
  }

  deleteTeam(id: string): Observable<void> {
    return this.http.delete<void>(this.url(`/admin/teams/${id}`));
  }

  createTrack(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/tracks'), payload);
  }

  updateTrack(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/tracks'), payload);
  }

  deleteTrack(id: string): Observable<void> {
    return this.http.delete<void>(this.url(`/admin/tracks/${id}`));
  }

  createRace(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/races'), payload);
  }

  updateRace(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/races'), payload);
  }

  deleteRace(id: string): Observable<void> {
    return this.http.delete<void>(this.url(`/admin/races/${id}`));
  }

  createCarParticipant(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/car-participants'), payload);
  }

  updateCarParticipant(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/car-participants'), payload);
  }

  deleteCarParticipant(id: string): Observable<void> {
    return this.http.delete<void>(this.url(`/admin/car-participants/${id}`));
  }

  updateCarParticipantDrivers(payload: unknown): Observable<unknown> {
    return this.http.patch(this.url('/admin/car-participants/drivers'), payload);
  }

  upsertRaceResult(payload: unknown): Observable<unknown> {
    return this.http.post(this.url('/admin/race-results'), payload);
  }

  updateUserRole(payload: { id: string; role: Role }): Observable<unknown> {
    return this.http.patch(this.url('/admin/users/role'), payload);
  }

  getAdminCars(query = ''): Observable<Car[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Car[]>(this.url('/admin/cars'), { params });
  }

  getAdminCarParticipants(query = ''): Observable<CarParticipant[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<CarParticipant[]>(this.url('/admin/car-participants'), { params });
  }

  getAdminChampionships(query = ''): Observable<Championship[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Championship[]>(this.url('/admin/championships'), { params });
  }

  getAdminRaces(query = ''): Observable<Race[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Race[]>(this.url('/admin/races'), { params });
  }

  getAdminTracks(query = ''): Observable<Track[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<Track[]>(this.url('/admin/tracks'), { params });
  }

  getAdminUsers(query = ''): Observable<AdminUser[]> {
    const params = query ? new HttpParams().set('q', query) : undefined;
    return this.http.get<AdminUser[]>(this.url('/admin/users'), { params });
  }

  private url(path: string): string {
    return `${this.baseUrl}${path}`;
  }
}
