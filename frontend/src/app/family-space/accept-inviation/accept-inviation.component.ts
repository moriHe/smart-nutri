import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { InvitationsService } from 'api/invitations/invitations.service';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-accept-inviation',
  templateUrl: './accept-inviation.component.html',
  styleUrls: ['./accept-inviation.component.css']
})
export class AcceptInviationComponent {
  isInitialized = false
  invitationSuccesful = false
  errorMessage = ""
  ngOnInit(): void {
    this.route.queryParams.subscribe(params => {
      const token = params["token"]
      if (!token) {
        this.errorMessage = "INVALID_TOKEN"
        this.isInitialized = true
      } else {
        this.invitationsService.acceptInvitation(token).subscribe(response => {
          if (response) {
            this.invitationSuccesful = true
            this.userService.getUser().subscribe()
          } else {
            this.errorMessage = "INVITATION_FAILED"
          }
          this.isInitialized = true
        }) 
      }
    })
  }

  constructor(
    private route: ActivatedRoute,
    private invitationsService: InvitationsService,
    private userService: UserService
    ) {}

}
