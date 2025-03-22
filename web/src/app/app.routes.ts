import { Routes } from '@angular/router';

import { HomeComponent } from './screens/home/home.component';
import { CatalogoComponent } from './screens/catalogo/catalogo.component';

export const routes: Routes = [
    {
        path: '',
        component: HomeComponent
    },
    {
        path: 'catalogo',
        component: CatalogoComponent
    }
];
