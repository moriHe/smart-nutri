import { Component, ChangeDetectorRef} from '@angular/core';
import { UserService } from 'api/user/user.service';
import { BehaviorSubject, finalize, take } from 'rxjs';
import {Auth, authState} from '@angular/fire/auth'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Smart Nutri';
  isInitialized = false
  
  // might not needed. seems to be working without
  authSubscription = authState(this.auth).pipe(
    take(1),
    finalize(() => {
      this.isInitialized = true
      this.cdr.detectChanges()
    })).subscribe()

  ngOnDestroy(): void {
    this.authSubscription.unsubscribe()
  }

  constructor(
    public userService: UserService, 
    private auth: Auth,
    private cdr: ChangeDetectorRef
    ) {}
}
