import {Component, OnDestroy, OnInit, ViewChild, ViewChildren} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Subject, takeUntil} from "rxjs";
import {EventService} from "../../shared/service/event.service";
import {filter} from "rxjs/operators";
import {EventEnum} from "../../shared/service/model/event-enum";
import {Wallet} from "../../client/model/response";
import {WalletService} from "../../client/service/wallet.service";
import {LineChartComponent} from "./line-chart/line-chart.component";

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
            this.loadWallet()
        })
    }

    ngOnInit(): void {
    }

    loadWallet() {
        this.walletService.getWalletCollection()
            .subscribe((wallets) => {
                const wallet = wallets.find(w => w.id == this.walletId)

                this.currentWallet = wallet;

                this.currentWallet.transactions = this.currentWallet.transactions.sort((t1, t2) => {
                        const date1 = new Date(t1.time)
                        const date2 = new Date(t2.time)
                        return date2.getTime() - date1.getTime()
                    }
                );

                this.walletCurrency = wallet.fiat;
            })
    }


    ngOnDestroy(): void {
        this.destroy$.next();
        this.destroy$.complete();
    }

}
