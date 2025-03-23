import { Component } from '@angular/core';

import { AnimalListItemInterface } from '../../../interface/animal-list-item';

@Component({
  selector: 'app-animal-list-item',
  templateUrl: './animal-list-item.component.html',
  styleUrl: './animal-list-item.component.css'
})

export class AnimalListItemComponent {
  animal: AnimalListItemInterface = {
    nome: 'Lulinha',
    especie: 'Cachorro',
    castrado: false,
    vacinado: true,
    raca: 'Ruim',
    porte: 'Medio',
    sexo: 0
  }
}
