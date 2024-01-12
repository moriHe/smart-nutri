import { Component, Inject, Input } from '@angular/core';
import { Router } from '@angular/router';
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

  onLogout() {
    this.userService.logout()
  }

  constructor(
    private router: Router,
    private userService: UserService
  ) {}
}