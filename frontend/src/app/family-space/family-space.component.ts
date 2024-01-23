import { Component } from '@angular/core';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-family-space',
  templateUrl: './family-space.component.html',
  styleUrls: ['./family-space.component.css']
})
export class FamilySpaceComponent {
  // TODO add update user -> activeFamilyId
  // TODO display familys in frontend
  // TODO display activeFamily separated on top
  // TODO if activeFamily role of user is OWNER, he/she can generate the link
  // TODO display the link with a copy button
  // TODO let user switch family space (dont forget to refetch user)
  ngOnInit(): void {
    this.userService.user$.subscribe(user => console.log(user)) 
  }

  constructor(private userService: UserService) {}
}
