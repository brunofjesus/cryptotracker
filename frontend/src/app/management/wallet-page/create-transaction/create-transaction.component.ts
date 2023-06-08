import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {UntypedFormBuilder, UntypedFormGroup, Validators} from "@angular/forms";
import {MessageService} from "primeng/api";
import {EventService} from "../../../shared/service/event.service";
import {EventEnum} from "../../../shared/service/model/event-enum";
import {WalletService} from "../../../client/service/wallet.service";
import {Transaction} from "../../../client/model/response";
import {CreateTransactionRequest} from "../../../client/model/request";

@Component({
    selector: 'app-create-transaction',
    templateUrl: './create-transaction.component.html',
    styleUrls: ['./create-transaction.component.scss']
})
export class CreateTransactionComponent implements OnInit {

    @Input()
    walletId: number;

    @Input()
    currency: string;

    @Output()
    onClose = new EventEmitter<void>();

    form: UntypedFormGroup;

    constructor(
        private fb: UntypedFormBuilder,
        private walletService: WalletService,
        private messageService: MessageService,
        private eventService: EventService
    ) {

        this.form = this.fb.group({
                datetime: [new Date(), [Validators.required]],
                coinValue: [undefined, Validators.required],
                coinAmount: [undefined, Validators.required],
                fiatInvested: [undefined, Validators.required],
                keepInserting: [false]
            }
        )
    }

    ngOnInit(): void {
    }

    submit() {
        const transaction: CreateTransactionRequest = {
            time: this.form.value['datetime'],
            cryptoValue: this.form.value['coinValue'].toString(),
            cryptoAmount: this.form.value['coinAmount'].toString(),
            fiatInvested: this.form.value['fiatInvested'].toString()
        }
        this.walletService.addTransaction(this.walletId, transaction)
            .subscribe({
                next: (transactionId) => {
                    this.messageService.add({
                        severity: 'success',
                        summary: 'Transaction created'
                    });

                    this.eventService.emitEvent({
                        type: EventEnum.TRANSACTION_CREATED,
                        value: transactionId
                    });

                    if (this.form.value["keepInserting"]) {
                        this.form.reset();
                    } else {
                        this.onClose.emit();
                    }
                }
            });
    }
}
