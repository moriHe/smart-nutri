import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-base-button',
  templateUrl: './base-button.component.html',
  styleUrls: ['./base-button.component.css']
})
export class BaseButtonComponent {
  @Input() onClickBaseBaseButton: () => void = () => {}

  onClick() {
    this.onClickBaseBaseButton()
  }
}
