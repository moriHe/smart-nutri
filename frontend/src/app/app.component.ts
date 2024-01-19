import { Component, ChangeDetectorRef} from '@angular/core';
import { UserService } from 'api/user/user.service';
import { SupabaseService } from 'api/supabase.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Smart Nutri';
  isInitialized = false
  
 

    ngOnInit(): void {
      this.supabaseService.authChanges((_, session) => {
        this.supabaseService.setSession(session)
        this.isInitialized = true
      })
    }



  constructor(
    private supabaseService: SupabaseService,
    public userService: UserService, 
    private cdr: ChangeDetectorRef
    ) {}
}
