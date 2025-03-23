import { Component } from '@angular/core';

import { IconButtonComponent } from '../../components/icon-button/icon-button.component';
import { AnimalListItemComponent } from '../../components/animal-list-item/animal-list-item.component';

@Component({
  selector: 'app-catalogo',
  imports: [IconButtonComponent ,AnimalListItemComponent],
  templateUrl: './catalogo.component.html',
  styleUrl: './catalogo.component.css'
})
export class CatalogoComponent {

}
