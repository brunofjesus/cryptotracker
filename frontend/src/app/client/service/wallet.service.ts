import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../../environments/environment";
import {Observable} from "rxjs";
import {Transaction, Wallet} from "../model/wallet";

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

    createWallet(param: { name: string; fiat: string; crypto: string }): Observable<number> {
        return this.http.post<number>(environment.apiUrl + "/wallet", param)
    }
}
