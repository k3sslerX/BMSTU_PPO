import { JsonPipe } from '@angular/common';
import { Component, OnDestroy, OnInit, computed, inject, signal } from '@angular/core';
import { NonNullableFormBuilder, ReactiveFormsModule } from '@angular/forms';
import { Observable, Subscription, distinctUntilChanged, finalize, forkJoin } from 'rxjs';
import { AdminUser, Car, CarParticipant, Championship, Driver, Race, Role, Team, Track } from '../../core/models/api.models';
import { ApiService } from '../../core/services/api.service';
import { AuthService } from '../../core/services/auth.service';
import { readApiError } from '../../core/utils/api-error';
import { AdminLookupComboComponent, LookupOption } from './admin-lookup-combo.component';

interface Feedback {
  kind: 'success' | 'error';
  title: string;
  text: string;
}

@Component({
  selector: 'app-admin-page',
  standalone: true,
  imports: [AdminLookupComboComponent, JsonPipe, ReactiveFormsModule],
  templateUrl: './admin-page.component.html',
  styleUrl: './admin-page.component.scss'
})
export class AdminPageComponent implements OnInit, OnDestroy {
  private readonly fb = inject(NonNullableFormBuilder);
  private readonly api = inject(ApiService);
  private readonly auth = inject(AuthService);
  private readonly autofillSubscriptions = new Subscription();

  protected readonly busyAction = signal<string | null>(null);
  protected readonly referencesBusy = signal(false);
  protected readonly feedback = signal<Feedback | null>(null);
  protected readonly lastResponse = signal<unknown | null>(null);
  protected readonly drivers = signal<Driver[]>([]);
  protected readonly teams = signal<Team[]>([]);
  protected readonly tracks = signal<Track[]>([]);
  protected readonly races = signal<Race[]>([]);
  protected readonly cars = signal<Car[]>([]);
  protected readonly carParticipants = signal<CarParticipant[]>([]);
  protected readonly championships = signal<Championship[]>([]);
  protected readonly users = signal<AdminUser[]>([]);
  protected readonly driverOptions = computed<LookupOption[]>(() =>
    this.drivers().map((driver) => ({
      id: driver.id,
      label: driver.name,
      subtitle: `${driver.nationality} • ${formatDate(driver.birthday)}`,
      searchText: driver.birthday
    }))
  );
  protected readonly teamOptions = computed<LookupOption[]>(() =>
    this.teams().map((team) => ({
      id: team.id,
      label: team.name,
      subtitle: team.country
    }))
  );
  protected readonly trackOptions = computed<LookupOption[]>(() =>
    this.tracks().map((track) => ({
      id: track.id,
      label: track.name,
      subtitle: `${track.country} • ${track.length} м • ${track.turns} поворотов`
    }))
  );
  protected readonly raceOptions = computed<LookupOption[]>(() =>
    this.races().map((race) => ({
      id: race.id,
      label: this.raceLabel(race),
      subtitle: `Type ${race.type} • Duration ${race.duration}`
    }))
  );
  protected readonly carOptions = computed<LookupOption[]>(() =>
    this.cars().map((car) => ({
      id: car.id,
      label: this.carLabel(car),
      subtitle: [`Год ${car.year}`, car.raceclass ? `Класс ${car.raceclass}` : ''].filter(Boolean).join(' • '),
      searchText: [car.year, car.raceclass].filter(Boolean).join(' ')
    }))
  );
  protected readonly carParticipantOptions = computed<LookupOption[]>(() =>
    this.carParticipants().map((participant) => ({
      id: participant.id,
      label: this.carParticipantLabel(participant),
      subtitle: this.carParticipantSubtitle(participant),
      searchText: [participant.number, participant.car_id, participant.team_id].join(' ')
    }))
  );
  protected readonly championshipOptions = computed<LookupOption[]>(() =>
    this.championships().map((championship) => ({
      id: championship.id,
      label: this.championshipLabel(championship),
      subtitle: championship.organizer,
      searchText: [championship.year, championship.organizer].filter(Boolean).join(' ')
    }))
  );
  protected readonly userOptions = computed<LookupOption[]>(() =>
    this.users()
      .filter((user) => user.id !== this.auth.userId())
      .map((user) => ({
        id: user.id,
        label: user.email,
        subtitle: `${user.name || 'без имени'} • ${user.role}`,
        searchText: user.name
      }))
  );
  protected readonly roleLookupOptions: LookupOption[] = [
    {
      id: 'user',
      label: 'user',
      subtitle: 'Обычный пользователь'
    },
    {
      id: 'admin',
      label: 'admin',
      subtitle: 'Администратор'
    }
  ];

  protected readonly driverForm = this.fb.group({
    id: [''],
    name: [''],
    nationality: [''],
    birthday: ['']
  });

  protected readonly teamForm = this.fb.group({
    id: [''],
    name: [''],
    country: ['']
  });

  protected readonly trackForm = this.fb.group({
    id: [''],
    name: [''],
    country: [''],
    length: [''],
    turns: ['']
  });

  protected readonly raceForm = this.fb.group({
    id: [''],
    name: [''],
    date: [''],
    type: [''],
    duration: [''],
    trackId: [''],
    championshipId: ['']
  });

  protected readonly carParticipantForm = this.fb.group({
    id: [''],
    carId: [''],
    teamId: [''],
    number: [''],
    driverIds: [[] as string[]]
  });

  protected readonly raceResultForm = this.fb.group({
    raceId: [''],
    carParticipantId: [''],
    finishPos: [''],
    qualifyingPos: ['']
  });

  protected readonly userRoleForm = this.fb.group({
    id: [''],
    role: ['user' as Role]
  });

  ngOnInit(): void {
    this.setupAutofill();
    this.loadReferences();
  }

  ngOnDestroy(): void {
    this.autofillSubscriptions.unsubscribe();
  }

  private setupAutofill(): void {
    this.autofillSubscriptions.add(
      this.driverForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillDriver(id))
    );
    this.autofillSubscriptions.add(
      this.teamForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillTeam(id))
    );
    this.autofillSubscriptions.add(
      this.trackForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillTrack(id))
    );
    this.autofillSubscriptions.add(
      this.raceForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillRace(id))
    );
    this.autofillSubscriptions.add(
      this.carParticipantForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillCarParticipant(id))
    );
    this.autofillSubscriptions.add(
      this.userRoleForm.controls.id.valueChanges.pipe(distinctUntilChanged()).subscribe((id) => this.autofillUserRole(id))
    );
  }

  private autofillDriver(id: string): void {
    const driver = findById(this.drivers(), id);
    if (!driver) {
      return;
    }

    this.driverForm.patchValue(
      {
        name: driver.name,
        nationality: driver.nationality,
        birthday: formatDate(driver.birthday)
      },
      { emitEvent: false }
    );
  }

  private autofillTeam(id: string): void {
    const team = findById(this.teams(), id);
    if (!team) {
      return;
    }

    this.teamForm.patchValue(
      {
        name: team.name,
        country: team.country
      },
      { emitEvent: false }
    );
  }

  private autofillTrack(id: string): void {
    const track = findById(this.tracks(), id);
    if (!track) {
      return;
    }

    this.trackForm.patchValue(
      {
        name: track.name,
        country: track.country,
        length: String(track.length),
        turns: String(track.turns)
      },
      { emitEvent: false }
    );
  }

  private autofillRace(id: string): void {
    const race = findById(this.races(), id);
    if (!race) {
      return;
    }

    this.raceForm.patchValue(
      {
        name: race.name,
        date: formatDate(race.date),
        type: String(race.type),
        duration: String(race.duration),
        trackId: race.track?.id ?? '',
        championshipId: race.championship_id
      },
      { emitEvent: false }
    );
  }

  private autofillCarParticipant(id: string): void {
    const participant = findById(this.carParticipants(), id);
    if (!participant) {
      return;
    }

    this.carParticipantForm.patchValue(
      {
        carId: participant.car_id,
        teamId: participant.team_id,
        number: participant.number,
        driverIds: participant.drivers?.map((driver) => driver.id).filter(isNonEmptyString) ?? []
      },
      { emitEvent: false }
    );
  }

  private autofillUserRole(id: string): void {
    const user = findById(this.users(), id);
    if (!user) {
      return;
    }

    this.userRoleForm.patchValue(
      {
        role: user.role
      },
      { emitEvent: false }
    );
  }

  protected loadReferences(): void {
    this.referencesBusy.set(true);

    forkJoin({
      drivers: this.api.getDrivers(),
      teams: this.api.getTeams(),
      tracks: this.api.getAdminTracks(),
      races: this.api.getAdminRaces(),
      cars: this.api.getAdminCars(),
      carParticipants: this.api.getAdminCarParticipants(),
      championships: this.api.getAdminChampionships(),
      users: this.api.getAdminUsers()
    })
      .pipe(finalize(() => this.referencesBusy.set(false)))
      .subscribe({
        next: ({ drivers, teams, tracks, races, cars, carParticipants, championships, users }) => {
          this.drivers.set(drivers);
          this.teams.set(teams);
          this.tracks.set(tracks);
          this.races.set(races);
          this.cars.set(cars);
          this.carParticipants.set(carParticipants);
          this.championships.set(championships);
          this.users.set(users);
        },
        error: (error) => {
          this.feedback.set({
            kind: 'error',
            title: 'Не удалось загрузить справочники',
            text: readApiError(error)
          });
        }
      });
  }

  protected submitDriver(mode: 'create' | 'update'): void {
    try {
      const raw = this.driverForm.getRawValue();
      const payload = compact({
        id: mode === 'update' ? required(raw.id, 'Driver UUID') : undefined,
        name: trimmed(raw.name) || undefined,
        nationality: trimmed(raw.nationality) || undefined,
        birthday: trimmed(raw.birthday) || undefined
      });

      if (mode === 'create' && (!payload.name || !payload.nationality || !payload.birthday)) {
        throw new Error('Для создания пилота нужны name, nationality и birthday.');
      }

      this.run(
        `driver-${mode}`,
        mode === 'create' ? this.api.createDriver(payload) : this.api.updateDriver(payload),
        mode === 'create' ? 'Пилот создан' : 'Пилот обновлён'
      );
    } catch (error) {
      this.setLocalError(error, 'Driver form');
    }
  }

  protected submitTeam(mode: 'create' | 'update'): void {
    try {
      const raw = this.teamForm.getRawValue();
      const payload = compact({
        id: mode === 'update' ? required(raw.id, 'Team UUID') : undefined,
        name: trimmed(raw.name) || undefined,
        country: trimmed(raw.country) || undefined
      });

      if (mode === 'create' && (!payload.name || !payload.country)) {
        throw new Error('Для создания команды нужны name и country.');
      }

      this.run(
        `team-${mode}`,
        mode === 'create' ? this.api.createTeam(payload) : this.api.updateTeam(payload),
        mode === 'create' ? 'Команда создана' : 'Команда обновлена'
      );
    } catch (error) {
      this.setLocalError(error, 'Team form');
    }
  }

  protected submitTrack(mode: 'create' | 'update'): void {
    try {
      const raw = this.trackForm.getRawValue();
      const payload = compact({
        id: mode === 'update' ? required(raw.id, 'Track UUID') : undefined,
        name: trimmed(raw.name) || undefined,
        country: trimmed(raw.country) || undefined,
        length: optionalNumber(raw.length),
        turns: optionalNumber(raw.turns)
      });

      if (mode === 'create' && (!payload.name || !payload.country || !payload.length || !payload.turns)) {
        throw new Error('Для создания трассы нужны name, country, length и turns.');
      }

      this.run(
        `track-${mode}`,
        mode === 'create' ? this.api.createTrack(payload) : this.api.updateTrack(payload),
        mode === 'create' ? 'Трасса создана' : 'Трасса обновлена'
      );
    } catch (error) {
      this.setLocalError(error, 'Track form');
    }
  }

  protected submitRace(mode: 'create' | 'update'): void {
    try {
      const raw = this.raceForm.getRawValue();
      const trackId = trimmed(raw.trackId);
      const payload = compact({
        id: mode === 'update' ? required(raw.id, 'Race UUID') : undefined,
        name: trimmed(raw.name) || undefined,
        date: trimmed(raw.date) || undefined,
        type: optionalNumber(raw.type),
        duration: optionalNumber(raw.duration),
        championship_id: trimmed(raw.championshipId) || undefined,
        track: trackId ? { id: trackId } : undefined
      });

      if (
        mode === 'create' &&
        (!payload.name || !payload.date || !payload.type || !payload.duration || !payload.championship_id || !payload.track)
      ) {
        throw new Error('Для создания гонки нужны name, date, type, duration, trackId и championshipId.');
      }

      this.run(
        `race-${mode}`,
        mode === 'create' ? this.api.createRace(payload) : this.api.updateRace(payload),
        mode === 'create' ? 'Гонка создана' : 'Гонка обновлена'
      );
    } catch (error) {
      this.setLocalError(error, 'Race form');
    }
  }

  protected submitCarParticipant(mode: 'create' | 'update' | 'drivers'): void {
    try {
      const raw = this.carParticipantForm.getRawValue();

      const payload = compact({
        id: mode !== 'create' ? required(raw.id, 'Car participant UUID') : undefined,
        car_id: mode !== 'drivers' ? trimmed(raw.carId) || undefined : undefined,
        team_id: mode !== 'drivers' ? trimmed(raw.teamId) || undefined : undefined,
        number: mode !== 'drivers' ? trimmed(raw.number) || undefined : undefined,
        drivers: parseDriverRefs(raw.driverIds)
      });

      if (mode === 'create' && (!payload.car_id || !payload.team_id || !payload.number)) {
        throw new Error('Для создания car participant нужны carId, teamId и number.');
      }

      let request$: Observable<unknown>;
      let successTitle: string;

      if (mode === 'create') {
        request$ = this.api.createCarParticipant(payload);
        successTitle = 'Car participant создан';
      } else if (mode === 'update') {
        request$ = this.api.updateCarParticipant(payload);
        successTitle = 'Car participant обновлён';
      } else {
        request$ = this.api.updateCarParticipantDrivers(payload);
        successTitle = 'Состав пилотов обновлён';
      }

      this.run(`car-participant-${mode}`, request$, successTitle);
    } catch (error) {
      this.setLocalError(error, 'Car participant form');
    }
  }

  protected submitRaceResult(): void {
    try {
      const raw = this.raceResultForm.getRawValue();
      const payload = {
        race_id: required(raw.raceId, 'Race UUID'),
        car_participant_id: required(raw.carParticipantId, 'Car participant UUID'),
        finish_pos: positiveNumber(raw.finishPos, 'Finish position'),
        qualifying_pos: positiveNumber(raw.qualifyingPos, 'Qualifying position')
      };

      this.run('race-result', this.api.upsertRaceResult(payload), 'Результат гонки записан');
    } catch (error) {
      this.setLocalError(error, 'Race result form');
    }
  }

  protected submitUserRole(): void {
    try {
      const raw = this.userRoleForm.getRawValue();
      const userId = required(raw.id, 'User UUID');
      if (userId === this.auth.userId()) {
        throw new Error('Нельзя менять роль самому себе.');
      }

      this.run(
        'user-role',
        this.api.updateUserRole({
          id: userId,
          role: roleValue(raw.role)
        }),
        'Роль пользователя обновлена'
      );
    } catch (error) {
      this.setLocalError(error, 'User role form');
    }
  }

  protected deleteDriver(): void {
    this.deleteSelected(
      'driver-delete',
      this.driverForm.controls.id.value,
      'пилота',
      'Пилот удалён',
      (id) => this.api.deleteDriver(id),
      () => this.driverForm.controls.id.setValue('')
    );
  }

  protected deleteTeam(): void {
    this.deleteSelected(
      'team-delete',
      this.teamForm.controls.id.value,
      'команду',
      'Команда удалена',
      (id) => this.api.deleteTeam(id),
      () => this.teamForm.controls.id.setValue('')
    );
  }

  protected deleteTrack(): void {
    this.deleteSelected(
      'track-delete',
      this.trackForm.controls.id.value,
      'трассу',
      'Трасса удалена',
      (id) => this.api.deleteTrack(id),
      () => this.trackForm.controls.id.setValue('')
    );
  }

  protected deleteRace(): void {
    this.deleteSelected(
      'race-delete',
      this.raceForm.controls.id.value,
      'гонку',
      'Гонка удалена',
      (id) => this.api.deleteRace(id),
      () => this.raceForm.controls.id.setValue('')
    );
  }

  protected deleteCarParticipant(): void {
    this.deleteSelected(
      'car-participant-delete',
      this.carParticipantForm.controls.id.value,
      'участника',
      'Участник удалён',
      (id) => this.api.deleteCarParticipant(id),
      () => this.carParticipantForm.controls.id.setValue('')
    );
  }

  private run(action: string, request$: Observable<unknown>, successTitle: string, onSuccess?: () => void): void {
    this.busyAction.set(action);
    this.feedback.set(null);

    request$
      .pipe(finalize(() => this.busyAction.set(null)))
      .subscribe({
        next: (response) => {
          onSuccess?.();
          this.lastResponse.set(response ?? { ok: true });
          this.feedback.set({
            kind: 'success',
            title: successTitle,
            text: 'Сервер принял запрос. Последний ответ можно посмотреть ниже.'
          });
          this.loadReferences();
        },
        error: (error) => {
          this.lastResponse.set(null);
          this.feedback.set({
            kind: 'error',
            title: 'Запрос отклонён',
            text: readApiError(error)
          });
        }
      });
  }

  private deleteSelected(
    action: string,
    rawId: unknown,
    entityName: string,
    successTitle: string,
    requestFactory: (id: string) => Observable<unknown>,
    resetSelection: () => void
  ): void {
    try {
      const id = required(rawId, 'UUID для удаления');
      if (!window.confirm(`Удалить ${entityName}? Это действие нельзя отменить.`)) {
        return;
      }

      this.run(action, requestFactory(id), successTitle, resetSelection);
    } catch (error) {
      this.setLocalError(error, 'Delete form');
    }
  }

  private setLocalError(error: unknown, title: string): void {
    this.feedback.set({
      kind: 'error',
      title,
      text: readApiError(error)
    });
  }

  protected raceLabel(race: Race): string {
    return `${race.name} · ${formatDate(race.date)}`;
  }

  protected carLabel(car: Car): string {
    return car.model;
  }

  protected carParticipantLabel(participant: CarParticipant): string {
    const team = this.entityName(this.teams(), participant.team_id);
    const car = this.entityName(this.cars(), participant.car_id, (item) => this.carLabel(item));
    return [`#${participant.number || 'без номера'}`, team, car].filter(Boolean).join(' · ');
  }

  protected carParticipantSubtitle(participant: CarParticipant): string {
    return participant.drivers?.map((driver) => driver.name).filter(Boolean).join(' • ') ?? '';
  }

  protected championshipLabel(championship: Championship): string {
    return `Season ${championship.year}`;
  }

  protected userLabel(user: AdminUser): string {
    return `${user.email} · ${user.name || 'без имени'} · ${user.role}`;
  }

  protected driverSelected(): boolean {
    return hasValue(this.driverForm.controls.id.value);
  }

  protected teamSelected(): boolean {
    return hasValue(this.teamForm.controls.id.value);
  }

  protected trackSelected(): boolean {
    return hasValue(this.trackForm.controls.id.value);
  }

  protected raceSelected(): boolean {
    return hasValue(this.raceForm.controls.id.value);
  }

  protected carParticipantSelected(): boolean {
    return hasValue(this.carParticipantForm.controls.id.value);
  }

  private entityName<T extends { id: string; name?: string }>(items: T[], id: string, label?: (item: T) => string): string {
    const item = items.find((value) => value.id === id);
    return item ? label?.(item) ?? item.name ?? '' : '';
  }
}

function required(value: unknown, label: string): string {
  const text = trimmed(value);
  if (!text) {
    throw new Error(`${label} обязателен.`);
  }

  return text;
}

function optionalNumber(value: unknown): number | undefined {
  const text = trimmed(value);
  if (!text) {
    return undefined;
  }

  const parsed = Number(text);
  if (Number.isNaN(parsed)) {
    throw new Error(`Числовое поле "${value}" содержит некорректное значение.`);
  }

  return parsed;
}

function positiveNumber(value: unknown, label: string): number {
  const parsed = optionalNumber(value);
  if (!parsed || parsed <= 0) {
    throw new Error(`${label} должен быть положительным числом.`);
  }

  return parsed;
}

function parseDriverRefs(value: unknown): Array<{ id: string }> | undefined {
  const values = Array.isArray(value) ? value : String(value ?? '').split(',');
  const ids = values.map((item) => trimmed(item)).filter(Boolean).map((id) => ({ id }));

  return ids.length ? ids : undefined;
}

function roleValue(value: unknown): Role {
  const role = trimmed(value);
  if (role !== 'admin' && role !== 'user') {
    throw new Error('Role должен быть user или admin.');
  }

  return role;
}

function hasValue(value: unknown): boolean {
  return trimmed(value).length > 0;
}

function findById<T extends { id: string }>(items: T[], id: string): T | undefined {
  return items.find((item) => item.id === id);
}

function isNonEmptyString(value: unknown): value is string {
  return typeof value === 'string' && value.trim().length > 0;
}

function trimmed(value: unknown): string {
  if (Array.isArray(value)) {
    return value.map((item) => trimmed(item)).filter(Boolean).join(',');
  }

  return String(value ?? '').trim();
}

function compact<T extends Record<string, unknown>>(value: T): T {
  return Object.fromEntries(
    Object.entries(value).filter(([, item]) => item !== undefined && item !== '')
  ) as T;
}

function formatDate(value: string): string {
  return value ? value.slice(0, 10) : 'без даты';
}
