import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FamilysService } from 'api/familys/familys.service';
import { UserService } from 'api/user/user.service';
import { take } from 'rxjs';


@Component({
  selector: 'app-signup-success',
  templateUrl: './signup-success.component.html',
  styleUrls: ['./signup-success.component.css']
})
export class SignupSuccessComponent {
  name = ""

  createFamily() {
    if (!this.name) {
      return
    }
    this.familysService.createFamily(this.name).subscribe(
      {
        next: () => {
          this.userService.getUser().subscribe(
            {next: () => {
              this.router.navigateByUrl("meine-rezepte")
            }}
          )
        }
    })
  }

  ngOnInit(): void {
    this.userService.user$.subscribe(user => {
      if (user) {
        return
      }
      this.userService.addUser().subscribe(() => {
        this.userService.getUser().subscribe()
      })
    })
  }

  constructor(private userService: UserService, private familysService: FamilysService, private router: Router) {}
}


