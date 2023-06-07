import {RouterModule} from '@angular/router';
import {NgModule} from '@angular/core';
import {ManagementComponent} from "./management.component";
import {DashboardComponent} from "./dashboard/dashboard.component";
import {WalletPageComponent} from "./wallet-page/wallet-page.component";
import {CreateWalletComponent} from "./create-wallet/create-wallet.component";

@NgModule({
    imports: [
        RouterModule.forChild([
            {
                path: '',
                component: ManagementComponent,
                children: [
                    {
                        path: '',
                        component: DashboardComponent,
                        pathMatch: 'full',
                    },
                    {
                        path: 'wallets/create',
                        component: CreateWalletComponent
                    },
                    {
                        path: 'wallets/:walletId',
                        component: WalletPageComponent
                    },
                    {
                        path: 'wallets/:walletId/edit',
                        component: CreateWalletComponent
                    }
                ]
            },
            {path: '**', redirectTo: '/notfound'},
        ])
    ],
    exports: [RouterModule]
})
export class ManagementRoutingModule {
}
