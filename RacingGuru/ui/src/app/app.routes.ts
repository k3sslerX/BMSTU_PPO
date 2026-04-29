import { Routes } from '@angular/router';
import { adminGuard } from './core/guards/admin.guard';
import { authGuard } from './core/guards/auth.guard';

export const routes: Routes = [
  {
    path: '',
    title: 'RacingGuru',
    loadComponent: () => import('./pages/home/home-page.component').then((m) => m.HomePageComponent)
  },
  {
    path: 'auth',
    title: 'RacingGuru | Auth',
    loadComponent: () => import('./pages/auth/auth-page.component').then((m) => m.AuthPageComponent)
  },
  {
    path: 'stats',
    pathMatch: 'full',
    redirectTo: 'stats/drivers'
  },
  {
    path: 'stats/drivers',
    title: 'RacingGuru | Driver Stats',
    loadComponent: () => import('./pages/stats/stats-page.component').then((m) => m.StatsPageComponent)
  },
  {
    path: 'stats/teams',
    title: 'RacingGuru | Team Stats',
    loadComponent: () => import('./pages/stats/team-stats-page.component').then((m) => m.TeamStatsPageComponent)
  },
  {
    path: 'sudoku',
    title: 'RacingGuru | Sudoku',
    loadComponent: () => import('./pages/sudoku/sudoku-page.component').then((m) => m.SudokuPageComponent)
  },
  {
    path: 'account',
    title: 'RacingGuru | Account',
    canActivate: [authGuard],
    loadComponent: () => import('./pages/account/account-page.component').then((m) => m.AccountPageComponent)
  },
  {
    path: 'admin',
    title: 'RacingGuru | Admin',
    canActivate: [adminGuard],
    loadComponent: () => import('./pages/admin/admin-page.component').then((m) => m.AdminPageComponent)
  },
  {
    path: '**',
    redirectTo: ''
  }
];
