export type Role = 'admin' | 'user';

export interface ApiErrorEnvelope {
  error?: {
    code?: string;
    message?: string;
  };
}

export interface TokenResponse {
  token: string;
}

export interface AuthUserResponse {
  name: string;
  email: string;
  role: Role;
}

export interface FavouritesResponse {
  favourite_drivers: Driver[];
  favourite_teams: Team[];
}

export interface Driver {
  id: string;
  name: string;
  birthday: string;
  nationality: string;
}

export interface Team {
  id: string;
  name: string;
  country: string;
}

export interface Car {
  id: string;
  model: string;
  year: number;
  raceclass?: string;
}

export interface Championship {
  id: string;
  year: number;
  organizer?: string;
}

export interface Track {
  id: string;
  name: string;
  country: string;
  length: number;
  turns: number;
}

export interface Race {
  id: string;
  name: string;
  date: string;
  type: number;
  duration: number;
  track: Partial<Track>;
  championship_id: string;
}

export interface CarParticipant {
  id: string;
  car_id: string;
  team_id: string;
  number: string;
  drivers: Partial<Driver>[];
}

export interface AdminUser {
  id: string;
  name: string;
  email: string;
  role: Role;
}

export interface RaceResult {
  race_id: string;
  car_participant_id: string;
  finish_pos: number;
  qualifying_pos: number;
}

export interface Stats {
  total_races: number;
  total_wins: number;
  total_podiums: number;
  total_points: number;
  total_poles: number;
  best_finish: number;
  best_qualifying: number;
  championships_wins: number;
  best_championship_position: number;
}

export interface DriverStats {
  driver: Driver;
  stats: Stats;
}

export interface TeamStats {
  team: Team;
  stats: Stats;
}

export interface DriverStatsResponse {
  driver_stats: DriverStats;
}

export interface TeamStatsResponse {
  team_stats: TeamStats;
}

export type ConditionOperator = 'gt' | 'lt' | 'gte' | 'lte' | 'eq';

export interface SudokuCondition {
  label: string;
  field: string;
  op: ConditionOperator;
  value: number;
}

export interface MatrixDrivers {
  field: Driver[][][];
  condition_specs: SudokuCondition[][];
}

export interface MatrixTeams {
  field: Team[][][];
  condition_specs: SudokuCondition[][];
}

export interface JwtClaims {
  user_id: string;
  role: Role;
  exp?: number;
  iat?: number;
}
