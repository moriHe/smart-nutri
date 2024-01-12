import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Component } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  isMobile = false

  ngOnInit(): void {
    this.breakpointObserver.observe([
      Breakpoints.Handset,
      Breakpoints.Tablet,
    ]).subscribe(result => {
      this.isMobile = result.matches;
    });
  }

  constructor(
    private breakpointObserver: BreakpointObserver
    ) {}
}

