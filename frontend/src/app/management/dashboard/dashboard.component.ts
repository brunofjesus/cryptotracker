import {Component, OnInit} from '@angular/core';
import {MenuItem} from 'primeng/api';
import {DashboardChartConfiguration} from "./dashboardChartConfiguration";
import {Wallet} from "../../client/model/response";
import {WalletService} from "../../client/service/wallet.service";

@Component({
    templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {

    chartColors = DashboardChartConfiguration.pieChartColors;
    barOptions = DashboardChartConfiguration.barChartOptions;

    items: MenuItem[];

    wallets: Wallet[] = [];

    investedData: any;
    currentValueData: any;
    pieOptions: any;

    barData: any;

    summary: any;

    constructor(
        private walletService: WalletService
    ) {
        walletService.getWalletCollection().subscribe(wallets => {
            this.wallets = wallets;
            this.initSummary();
            this.initCharts();
        });
    }

    ngOnInit() {
    }

    private initSummary() {
        let summary = [];

        this.wallets.forEach(w => {
            let currencySummary = summary.find(s => s.fiatCurrency == w.fiat)
            if (!currencySummary) {
                currencySummary = {
                    fiatCurrency: w.fiat,
                    invested: 0,
                    value: 0
                };
                summary.push(currencySummary);
            }

            currencySummary['invested'] += Number(w.totalFiatInvested)
            currencySummary['value'] += Number(w.currentFiatValue);
        })

        console.log("SUMMARY", summary)
        this.summary = summary;
    }

    private initCharts() {
        this.investedData = {
            labels: this.wallets.map(w => w.name),
            datasets: [
                {
                    data: this.wallets.map(w => w.totalFiatInvested),
                    ...this.chartColors
                }
            ]
        };

        this.currentValueData = {
            labels: this.wallets.map(w => w.name),
            datasets: [
                {
                    data: this.wallets.map(w => w.currentFiatValue),
                    ...this.chartColors
                }
            ]
        }

        this.pieOptions = {
            plugins: {
                legend: {
                    labels: {
                        fontColor: '#A0A7B5'
                    }
                }
            }
        };

        this.barData = {
            labels: this.wallets.map(w => w.name),
            datasets: [
                {
                    label: 'Coin return on investment',
                    backgroundColor: '#90cd93',
                    data: this.wallets
                        .map(w => parseFloat(w.totalFiatInvested) == 0 ? "0" : w.returnOnInvestmentPercent)
                }
            ]
        };
    }
}
