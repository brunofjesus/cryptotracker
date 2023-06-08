import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../../environments/environment";
import {Observable} from "rxjs";
import {Transaction, Wallet} from "../model/response";
import {CreateTransactionRequest, CreateWalletRequest} from "../model/request";

@Injectable({
    providedIn: 'root'
})
export class WalletService {

    constructor(
        private http: HttpClient,
    ) {
    }

    getWalletCollection(): Observable<Wallet[]> {
        return this.http.get<Wallet[]>(environment.apiUrl + "/wallet")
    }

    deleteTransaction(transaction: Transaction) {
        //TODO: not implemented yet
        return new Observable();
    }

    deleteWalletById(id: number) {
        //TODO: not implemented yet
        return new Observable();
    }

    createWallet(param: CreateWalletRequest): Observable<number> {
        return this.http.post<number>(environment.apiUrl + "/wallet", param)
    }

    editWallet(walletId: number, param: CreateWalletRequest): Observable<void> {
        return this.http.put<void>(environment.apiUrl + "/wallet/" + walletId, param)
    }

    addTransaction(walletId: number, param: CreateTransactionRequest): Observable<number> {
        return this.http.post<number>(environment.apiUrl + "/wallet/" + walletId + "/transaction", param)
    }
}
