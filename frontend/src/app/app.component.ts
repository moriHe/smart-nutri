import { Component } from '@angular/core';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Smart Nutri';
  isInitializedSubscription = this.userService.isInitialized$.subscribe()
  userSubscription = this.userService.user.subscribe()

  ngOnDestroy(): void {
    this.userSubscription.unsubscribe()
    this.isInitializedSubscription.unsubscribe()
  }

  constructor(public userService: UserService) {}
}
