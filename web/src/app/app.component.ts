import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

import { GaartFooterComponent } from './components/gaart-footer/gaart-footer.component';
import { GaartHeaderComponent } from './components/gaart-header/gaart-header.component';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, GaartHeaderComponent, GaartFooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'GAART';
}
