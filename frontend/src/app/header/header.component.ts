import { Component } from '@angular/core';
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

  constructor(private userService: UserService) {}
}
