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
export class WelcomeComponent {
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

  constructor(private userService: UserService, private familysService: FamilysService, private router: Router) {}
}


