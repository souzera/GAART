import { Component } from '@angular/core';

@Component({
  selector: 'app-gaart-footer',
  imports: [],
  templateUrl: './gaart-footer.component.html',
  styleUrl: './gaart-footer.component.css'
})

export class GaartFooterComponent {
  year = new Date().getFullYear();
}
