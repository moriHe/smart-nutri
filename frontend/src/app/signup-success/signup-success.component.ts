import { Component } from '@angular/core';
import { Auth, authState, User } from '@angular/fire/auth';
import { Subscription } from 'rxjs';


@Component({
  selector: 'app-signup-success',
  templateUrl: './signup-success.component.html',
  styleUrls: ['./signup-success.component.css']
})
export class SignupSuccessComponent {
  authState$ = authState(this.auth);
  authStateSubscription!: Subscription;
  

  ngOnInit(): void {
    this.authStateSubscription = this.authState$.subscribe((aUser: User | null) => {
      console.log('Current User:', aUser);
      if (aUser) {
        console.log('User is logged in');
      } else {
        console.log('User is not logged in');
      }
    });

    // You can also access the current user outside of the subscription
    const currentUser = this.auth.currentUser;
    if (currentUser) {
      console.log('Current User (outside subscription):', currentUser);
      console.log('User is logged in (outside subscription)');
    } else {
      console.log('User is not logged in (outside subscription)');
    }
  }

  ngOnDestroy(): void {
    // Unsubscribe from the authState observable to prevent memory leaks
    if (this.authStateSubscription) {
      this.authStateSubscription.unsubscribe();
    }
  }

  constructor(
    private auth: Auth
    ) {}
}
