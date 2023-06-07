import {Component, Input, OnInit} from '@angular/core';
import {ConfirmationService, MessageService} from "primeng/api";
import {EventService} from "../../../shared/service/event.service";
import {EventEnum} from "../../../shared/service/model/event-enum";
import {Transaction} from "../../../client/model/wallet";
import {WalletService} from "../../../client/service/wallet.service";

@Component({
    selector: 'app-transaction-table',
    templateUrl: './transaction-table.component.html',
    styleUrls: ['./transaction-table.component.scss']
})
export class TransactionTableComponent implements OnInit {

    @Input()
    walletId: number;

    @Input()
    transactions: Transaction[];

    @Input()
    currency: string;

    selectedItem: Transaction;
    displayCreateDialog: boolean = false;

    constructor(
        private confirmationService: ConfirmationService,
        private messageService: MessageService,
        private walletService: WalletService,
        private eventService: EventService
    ) {
    }

    ngOnInit(): void {
    }

    rowSelected() {
        console.log("Selected", this.selectedItem)
    }

    deleteTransaction(selectedItem: Transaction) {
        this.confirmationService.confirm({
            message: 'Do you want to delete this record?',
            header: 'Delete Confirmation',
            icon: 'pi pi-info-circle',
            accept: () => {
                this.walletService.deleteTransaction(selectedItem)
                    .subscribe({
                        next: () => {
                            this.messageService.add({severity: 'success', summary: 'Transaction deleted'});
                            this.eventService.emitEvent({
                                type: EventEnum.TRANSACTION_DELETED,
                                value: selectedItem
                            });
                        }
                    });
            }
        });
    }
}
