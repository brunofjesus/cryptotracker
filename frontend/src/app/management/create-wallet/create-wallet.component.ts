import {Component, OnDestroy, OnInit} from '@angular/core';
import {UntypedFormBuilder, UntypedFormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import {EventService} from "../../shared/service/event.service";
import {EventEnum} from "../../shared/service/model/event-enum";
import {Subject, takeUntil} from "rxjs";
import {Wallet} from "../../client/model/response";
import {WalletService} from "../../client/service/wallet.service";

@Component({
    selector: 'app-create-wallet',
    templateUrl: './create-wallet.component.html',
    styleUrls: ['./create-wallet.component.scss']
})
export class CreateWalletComponent implements OnInit, OnDestroy {

    protected destroy$ = new Subject<void>();

    walletInEdition: Wallet;

    form: UntypedFormGroup;

    constructor(
        private fb: UntypedFormBuilder,
        private walletService: WalletService,
        private router: Router,
        private route: ActivatedRoute,
        private eventService: EventService
    ) {
        this.route.params
            .pipe(takeUntil(this.destroy$))
            .subscribe((params) => {
                if (params.walletId) {
                    this.walletService.getWalletCollection()
                        .subscribe(wallets => {
                            const wallet = wallets.find(w => w.id == params.walletId)
                            this.walletInEdition = wallet;

                            this.form.patchValue({
                                id: wallet.id,
                                name: wallet.name,
                                coin: wallet.crypto,
                                fiatCurrency: wallet.fiat
                            })
                        });
                }
            });
    }

    ngOnInit(): void {
        this.form = this.fb.group({
            id: [],
            name: ['', Validators.required],
            coin: ['', Validators.required],
            fiatCurrency: ['EUR', Validators.required]
        });
    }

    submit() {
        const handler = (walletId) => {
            this.eventService.emitEvent({
                type: EventEnum.WALLET_CREATED,
                value: walletId
            });

            this.router.navigate(['/wallets', walletId]);
        };

        if (this.walletInEdition) {
            this.walletService.editWallet(
                this.form.value['id'],
                {
                    name: this.form.value['name'],
                    crypto: this.form.value['coin'],
                    fiat: this.form.value['fiatCurrency']
                }
            ).subscribe(handler);
        } else {
            this.walletService.createWallet({
                name: this.form.value['name'],
                crypto: this.form.value['coin'],
                fiat: this.form.value['fiatCurrency']
            }).subscribe(handler);
        }
    }

    ngOnDestroy(): void {
        this.destroy$.next();
        this.destroy$.complete();
    }
}
