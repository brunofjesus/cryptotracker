import {Component, OnDestroy, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Subject, takeUntil} from "rxjs";
import {EventService} from "../../shared/service/event.service";
import {filter} from "rxjs/operators";
import {EventEnum} from "../../shared/service/model/event-enum";
import {Wallet} from "../../client/model/wallet";
import {WalletService} from "../../client/service/wallet.service";

@Component({
    selector: 'app-wallet-page',
    templateUrl: './wallet-page.component.html',
    styleUrls: ['./wallet-page.component.scss']
})
export class WalletPageComponent implements OnInit, OnDestroy {

    protected destroy$ = new Subject<void>();

    walletId: number = null;
    walletCurrency: string;
    currentWallet: Wallet;
    constructor(
        private route: ActivatedRoute,
        private walletService: WalletService,
        private eventService: EventService
    ) {
        this.route.params
            .pipe(takeUntil(this.destroy$))
            .subscribe((params) => {
                this.walletId = params.walletId;
                this.loadWallet();
            });

        this.eventService.event$
            .pipe(filter(event =>
                event.type === EventEnum.TRANSACTION_CREATED ||
                event.type === EventEnum.TRANSACTION_DELETED
            )).subscribe((event) => {
            if (event.type === EventEnum.TRANSACTION_CREATED) {
                this.currentWallet.transactions.push(event.value);
            } else if (event.type === EventEnum.TRANSACTION_DELETED) {
                this.currentWallet.transactions.splice(
                    this.currentWallet.transactions.indexOf(event.value), 1
                );
            }

            this.currentWallet.transactions = this.currentWallet.transactions.sort((t1, t2) =>
                t2.time.getDate() - t1.time.getDate()
            );
        });
    }

    loadWallet() {
        this.walletService.getWalletCollection()
            .subscribe((wallets) => {
                const wallet = wallets.find(w => w.id == this.walletId)

                this.currentWallet = wallet;
                this.walletCurrency = wallet.fiat;
            })
    }

    ngOnInit(): void {
    }

    ngOnDestroy(): void {
        this.destroy$.next();
        this.destroy$.complete();
    }

}
