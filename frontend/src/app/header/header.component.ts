import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {

  onLogout() {
    this.userService.logout()
  }

  switchScreen(url: string) {
    this.router.navigate([url])
  }

  isScreenOpen(url: string) {
    if (this.router.url === url) {
      return true
    }
    return false
  }

  constructor(
    private userService: UserService,
    private router: Router
    ) {}
}
