import {Component, Input, OnInit} from '@angular/core';
import {DatePipe} from "@angular/common";
import {Transaction} from "../../../client/model/response";

@Component({
    selector: 'app-line-chart',
    templateUrl: './line-chart.component.html',
    styleUrls: ['./line-chart.component.scss']
})
export class LineChartComponent implements OnInit {

    @Input()
    transactions: Transaction[]

    lineData: any;
    lineOptions: any;

    constructor(
        private datePipe: DatePipe
    ) {
        this.lineOptions = {
            plugins: {
                legend: {
                    labels: {
                        fontColor: '#A0A7B5'
                    }
                }
            },
            scales: {
                x: {
                    ticks: {
                        color: '#A0A7B5'
                    },
                    grid: {
                        color: 'rgba(160, 167, 181, .3)',
                    }
                },
                y: {
                    ticks: {
                        color: '#A0A7B5'
                    },
                    grid: {
                        color: 'rgba(160, 167, 181, .3)',
                    }
                },
            }
        };
    }

    ngOnInit(): void {
        let previousTransactionLen = 0
        setInterval(() => {
            if (this.transactions.length != previousTransactionLen) {
                previousTransactionLen = this.transactions.length
                this.load()
            }
        }, 100)
    }

    load(): void {
        const reversedChart = [...this.transactions].reverse();

        this.lineData = {
            labels: reversedChart.map(m => m.time)
                .map(m => this.datePipe.transform(m, 'yyyy-MM-dd HH:mm')),
            datasets: [
                {
                    label: 'Invested Sum',
                    data: reversedChart.map(m => m.totalFiatInvested),
                    fill: false,
                    backgroundColor: '#f9ae61',
                    borderColor: '#f9ae61',
                    tension: .4
                },
                {
                    label: 'Wallet Value',
                    data: reversedChart.map(m => m.fiatWalletValue),
                    fill: false,
                    backgroundColor: '#90cd93',
                    borderColor: '#90cd93',
                    tension: .4
                }
            ]
        };
    }

}
