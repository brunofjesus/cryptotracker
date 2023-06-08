export interface Wallet {
    id: number
    name: string
    crypto: string
    fiat: string
    cryptoUnitValue: string
    totalCryptoAmount: string
    totalFiatInvested: string
    currentFiatValue: string
    returnOnInvestment: string
    returnOnInvestmentPercent: string
    cryptoUnitValueBreakEven: string
    transactions: Transaction[]
}

export interface Transaction {
    id: number
    walletId: number
    time: Date
    cryptoValue: string
    cryptoAmount: string
    fiatInvested: string
    fiatValue?: string
    totalCryptoAmount?: string
    totalFiatInvested?: string
    fiatWalletValue?: string
    fiatReturnOnInvestment?: string
}