export interface CreateWalletRequest {
    name: string
    fiat: string
    crypto: string
}

export interface CreateTransactionRequest {
    time: Date
    cryptoValue: string
    cryptoAmount: string
    fiatInvested: string
}