import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {DashboardComponent} from "./dashboard/dashboard.component";
import {ManagementComponent} from "./management.component";
import {ManagementRoutingModule} from "./management-routing.module";
import {SharedModule} from "../shared/shared.module";
import {WalletPageComponent} from './wallet-page/wallet-page.component';
import {TransactionTableComponent} from './wallet-page/transaction-table/transaction-table.component';
import {CreateTransactionComponent} from './wallet-page/create-transaction/create-transaction.component';
import {ReactiveFormsModule} from "@angular/forms";
import { WalletPageHeaderComponent } from './wallet-page/wallet-page-header/wallet-page-header.component';
import { LineChartComponent } from './wallet-page/line-chart/line-chart.component';
import { CreateWalletComponent } from './create-wallet/create-wallet.component';


@NgModule({
    declarations: [
        ManagementComponent,
        DashboardComponent,
        WalletPageComponent,
        TransactionTableComponent,
        CreateTransactionComponent,
        WalletPageHeaderComponent,
        LineChartComponent,
        CreateWalletComponent,
    ],
    imports: [
        CommonModule,
        ManagementRoutingModule,
        ReactiveFormsModule,
        SharedModule
    ]
})
export class ManagementModule {
}
