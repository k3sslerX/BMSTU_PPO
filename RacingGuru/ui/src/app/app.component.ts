import { Component, OnInit, inject } from '@angular/core';
import { Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { ApiService } from './core/services/api.service';
import { AuthService } from './core/services/auth.service';

@Component({
  selector: 'app-root',
  imports: [RouterLink, RouterLinkActive, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  private readonly api = inject(ApiService);
  private readonly router = inject(Router);

  protected readonly auth = inject(AuthService);
  protected readonly year = new Date().getFullYear();

  ngOnInit(): void {
    this.loadProfile();
  }

  protected logout(): void {
    this.auth.logout();
    void this.router.navigate(['/auth']);
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
        // Header still works with the local token while the backend is restarting.
      }
    });
  }
}
