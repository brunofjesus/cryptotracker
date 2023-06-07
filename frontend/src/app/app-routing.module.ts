import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';
import {AppMainComponent} from './app.main.component';

@NgModule({
    imports: [
        RouterModule.forRoot([
            {
                path: '',
                component: AppMainComponent,
                children: [
                    {
                        path: '',
                        loadChildren: () => import('./management/management.module').then(m => m.ManagementModule),
                    },

                ]
            },
            {path: '**', redirectTo: 'pages/empty'},
        ], {scrollPositionRestoration: 'enabled'})
    ],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
