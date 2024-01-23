import { ChangeDetectionStrategy, ChangeDetectorRef, Component } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { InvitationsService } from 'api/invitations/invitations.service';
import { User, UserFamily, UserFamilys, UserService } from 'api/user/user.service';

@Component({
  selector: 'app-family-space',
  templateUrl: './family-space.component.html',
  styleUrls: ['./family-space.component.css']
})
export class FamilySpaceComponent {
  isLoading = true
  user: User | null = null
  activeFamily?: UserFamily
  familys: UserFamilys = []
  invitationLink: string = ""
  // TODO display familys in frontend
  // TODO display activeFamily separated on top
  // TODO if activeFamily role of user is OWNER, he/she can generate the link
  // TODO display the link with a copy button
  // TODO let user switch family space (dont forget to refetch user)

  switchFamily(familyId: number) {
    this.userService.updateUserFamily(familyId).subscribe((response) => {
      if (response) {
        this.userService.getUser().subscribe()
      }
    })
  }

  generateInvitationLink() {
    this.invitationsServide.getInvitationLink().subscribe((url) => {
      this.invitationLink = url
      this.copyToClipboard(url)
    })
  }

  copyToClipboard(url: string) {
    navigator.clipboard.writeText(url)
      .then(() => {
        this.snackbar.open("Link kopiert")
      })
      .catch((error) => console.log(error))
  }

  getData() {
    this.userService.user$.subscribe(user => {
      this.user = user
      this.userService.getUserFamilys().subscribe((familys) => {
        this.familys = familys.filter((family) => family.familyId != this.user?.activeFamilyId)
        this.activeFamily = familys.find((family) => family.familyId === this.user?.activeFamilyId)
        this.isLoading = false
      })
    }) 
  }
  ngOnInit(): void {
    this.getData()
  }

  constructor(
    private snackbar: MatSnackBar,
    private invitationsServide: InvitationsService,
    private userService: UserService
    ) {}
}
