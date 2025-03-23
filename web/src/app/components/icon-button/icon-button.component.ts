import { Component } from '@angular/core';

import { IconButtonInterface } from '../../../interface/icon-button';

@Component({
  selector: 'app-icon-button',
  imports: [],
  templateUrl: './icon-button.component.html',
  styleUrl: './icon-button.component.css'
})
export class IconButtonComponent {
  props: IconButtonInterface = {
    icon: 'https://img.freepik.com/vetores-premium/cute-avatar-beagle-cabeca-simples-desenho-animado-vetor-ilustracao-cao-racas-natureza-conceito-icone-isolado_772770-330.jpg',
    label: 'cachorro',
    onClick: () => {
      console.log('Icon Button Clicked');
    }
  }
}
