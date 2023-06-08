import {Component, Input, OnInit} from '@angular/core';
import {ConfirmationService, MenuItem} from "primeng/api";
import {Router} from "@angular/router";
import {EventService} from "../../../shared/service/event.service";
import {EventEnum} from "../../../shared/service/model/event-enum";
import {Wallet} from "../../../client/model/response";
import {WalletService} from "../../../client/service/wallet.service";

@Component({
    selector: 'app-wallet-page-header',
    templateUrl: './wallet-page-header.component.html',
    styleUrls: ['./wallet-page-header.component.scss']
})
export class WalletPageHeaderComponent implements OnInit {

    @Input()
    currentWallet: Wallet;

    @Input()
    transactionCount: number;

    walletCurrency: string;

    optionsMenu: MenuItem[];

    constructor(
        private confirmationService: ConfirmationService,
        private walletService: WalletService,
        private router: Router,
        private eventService: EventService
    ) {
    }

    ngOnInit(): void {
        this.walletCurrency = this.currentWallet.fiat;

        this.optionsMenu = [
            {
                label: 'Edit',
                icon: 'pi pi-fw pi-pencil',
                routerLink: '/wallets/' + this.currentWallet.id + '/edit'
            },
            {
                separator: true
            },
            {
                label: 'Delete', icon: 'pi pi-fw pi-trash', command: () => {
                    this.deleteWallet();
                }
            }
        ];
    }

    deleteWallet() {
        this.confirmationService.confirm({
            message: 'Do you want to delete this record?',
            header: 'Delete Confirmation',
            icon: 'pi pi-info-circle',
            accept: () => {
                this.walletService.deleteWalletById(this.currentWallet.id)
                    .subscribe(() => {
                        this.eventService.emitEvent({
                            type: EventEnum.WALLET_DELETED,
                            value: this.currentWallet
                        })
                        this.router.navigateByUrl("/")
                    })
            }
        });
    }
}
