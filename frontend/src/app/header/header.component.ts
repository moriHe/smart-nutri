import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Component, HostListener, Input, ViewChild } from '@angular/core';
import { SupabaseService } from 'api/supabase.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  isMobile = false
  isLoggedIn = false
  @Input() toggleIsMobileSidenavOpen!: () => void
  @Input() setIsMobileSidenavOpenFalse!: () => void
  @ViewChild('sidenav') sidenav: any;

  @HostListener("document:click", ["$event"])
  onDocumentClick(event: any): void {
    if (!this.sidenav) {
      return
    }
    if (this.sidenav._elementRef.nativeElement.contains(event.target)) {
      return
    }
    if (event.target.parentElement.getAttribute("aria-label") === "menu") {
      this.sidenav.toggle()
      this.toggleIsMobileSidenavOpen()
      return
    }
    this.closeSidenav()
  }

  closeSidenav = () => {
    this.sidenav.close()
    this.setIsMobileSidenavOpenFalse()
  }


  ngOnInit(): void {
    this.breakpointObserver.observe([
      Breakpoints.Handset,
      Breakpoints.Tablet,
    ]).subscribe(result => {
      this.isMobile = result.matches;
    });

    this.supabaseService.session$.subscribe((session) => {
      if (session) {
        this.isLoggedIn = true
      } else {
        this.isLoggedIn = false
      }
    })
  }

  constructor(
    private breakpointObserver: BreakpointObserver,
    private supabaseService: SupabaseService
    ) {}
}

