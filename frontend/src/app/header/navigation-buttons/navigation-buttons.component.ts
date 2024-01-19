import { Component, Inject, Input } from '@angular/core';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-navigation-buttons',
  templateUrl: './navigation-buttons.component.html',
  styleUrls: ['./navigation-buttons.component.css']
})
export class NavigationButtonsComponent {
  @Input() isMobile!: boolean

  switchScreen(url: string) {
    this.router.navigate([url])
  }

  isScreenOpen(url: string) {
    if (this.router.url === url) {
      return true
    }
    return false
  }

  getHighlighting(url: string) {
    const isOpen = this.router.url === url
    if (this.isMobile) {
      return {'bg-yellow-400': isOpen}
    }

    return {
      'bg-yellow-400': isOpen,
        'hover:bg-yellow-400/30': !isOpen
    }
  }

  async onLogout() {
    const {error} = await this.supabaseService.logout()
    console.log(error)
    this.supabaseService.setSession(null)
    this.router.navigate(["/"])
  }

  constructor(
    private router: Router,
    private userService: UserService,
    private supabaseService: SupabaseService
  ) {}
}
