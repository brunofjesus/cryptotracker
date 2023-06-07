export class DashboardChartConfiguration {

    static pieChartColors = {
        backgroundColor: [
            //300
            "#75bef8",
            "#90cd93",
            "#fdd87d",
            "#61d5e4",
            "#f1749e",
            "#8893d1",
            "#61beb5",
            "#f9ae61",
            "#9caeb7",
            "#c279ce",
            //200
            "#a0d2fa",
            "#b2ddb4",
            "#fde4a5",
            "#91e2ed",
            "#f69ebc",
            "#acb4df",
            "#91d2cc",
            "#fbc791",
            "#bbc7cd",
            "#d4a2dd",
            //100
            "#cae6fc",
            "#d4ecd5",
            "#fef0cd",
            "#c2eff5",
            "#fac9da",
            "#d1d5ed",
            "#c2e6e2",
            "#fde0c2",
            "#d9e0e3",
            "#e7cbec"
        ],
        hoverBackgroundColor: [
            //400,
            "#4baaf5",
            "#6ebe71",
            "#fccc55",
            "#30c9dc",
            "#ed4981",
            "#6372c3",
            "#30aa9f",
            "#f79530",
            "#7e96a1",
            "#af50bf",
            //300
            "#75bef8",
            "#90cd93",
            "#fdd87d",
            "#61d5e4",
            "#f1749e",
            "#8893d1",
            "#61beb5",
            "#f9ae61",
            "#9caeb7",
            "#c279ce",
            //200
            "#a0d2fa",
            "#b2ddb4",
            "#fde4a5",
            "#91e2ed",
            "#f69ebc",
            "#acb4df",
            "#91d2cc",
            "#fbc791",
            "#bbc7cd",
            "#d4a2dd"
        ]
    };
    static barChartOptions = {
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
                    color:  'rgba(160, 167, 181, .3)',
                }
            },
            y: {
                ticks: {
                    color: '#A0A7B5'
                },
                grid: {
                    color:  'rgba(160, 167, 181, .3)',
                }
            },
        }
    };
}