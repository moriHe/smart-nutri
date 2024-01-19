import { Component, ChangeDetectorRef} from '@angular/core';
import { UserService } from 'api/user/user.service';
import { SupabaseService } from 'api/supabase.service';
import { finalize } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Smart Nutri';
  isInitialized = this.supabaseService.isAppIntialized$
  
 





  constructor(
    private supabaseService: SupabaseService, 
    ) {}
}
