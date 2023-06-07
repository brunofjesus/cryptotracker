import {Component, OnInit} from '@angular/core';
import {MenuService} from "../shared/service/app.menu.service";
import {EventService} from "../shared/service/event.service";
import {filter} from "rxjs/operators";
import {EventEnum} from "../shared/service/model/event-enum";
import {WalletService} from "../client/service/wallet.service";


@Component({
    selector: 'app-manager',
    template: '<router-outlet></router-outlet>'
})
export class ManagementComponent implements OnInit {

    constructor(
        private walletService: WalletService,
        private menuService: MenuService,
        private eventService: EventService
    ) {
        this.eventService.event$
            .pipe(filter((event) =>
                [EventEnum.WALLET_CREATED, EventEnum.WALLET_UPDATED, EventEnum.WALLET_DELETED]
                    .map(e => e.toString())
                    .includes(event.type)))
            .subscribe((event) => {
                this.initializeMenu();
            });
    }

    ngOnInit(): void {
        this.initializeMenu();
    }

    initializeMenu() {
        this.walletService.getWalletCollection().subscribe(
            (wallets) => {
                this.menuService.setMenuItems([
                    {
                        label: 'Home',
                        items: [
                            {label: 'Dashboard', icon: 'pi pi-fw pi-home', routerLink: ['/']}
                        ]
                    },
                    {
                        label: 'Wallets',
                        items: [
                            ...wallets.map(w => {
                                return {
                                    label: w.name, coin: w.crypto.toLowerCase(), routerLink: ['/wallets/' + w.id]
                                }
                            }),
                            {
                                label: 'Add wallet', icon: 'pi pi-plus', routerLink: ['/wallets/create']
                            }
                        ]
                    }
                ]);
            }
        )
    }


}
