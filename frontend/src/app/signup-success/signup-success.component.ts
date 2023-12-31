import { Component } from '@angular/core';
import { UserService } from 'api/user/user.service';


@Component({
  selector: 'app-signup-success',
  templateUrl: './signup-success.component.html',
  styleUrls: ['./signup-success.component.css']
})
export class SignupSuccessComponent {

  // authState = this.userService.authState$.subscribe((value) => {
  //   console.log(value)
  // })
  
  constructor(private userService: UserService) {}
}
