import { Component } from '@angular/core';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent {

  deleteAccount() {
    this.userService.deleteUser().subscribe()
  }

  constructor(private userService: UserService) {}
}
