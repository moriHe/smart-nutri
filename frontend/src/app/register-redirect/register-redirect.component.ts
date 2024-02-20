import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-register-redirect',
  templateUrl: './register-redirect.component.html',
  styleUrls: ['./register-redirect.component.css']
})
export class RegisterRedirectComponent {

  ngOnInit(): void {
    this.userService.user$.subscribe(user => {
      if (user) {
         this.router.navigate(["/willkommen"])
      } else {
        this.userService.addUser().subscribe(() => {
          this.userService.getUser().subscribe((activeFamilyId) => {
            if (activeFamilyId) {
              return this.router.navigate(["/meine-rezepte"])
            } else {
              return this.router.navigate(["/willkommen"])
            }
          })
        })
      }
    })
  }

  constructor(
    private router: Router,
    private userService: UserService
    ) {}
}
