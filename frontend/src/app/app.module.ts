import {NgModule} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {BrowserModule} from '@angular/platform-browser';
import {LocationStrategy, HashLocationStrategy, DatePipe} from '@angular/common';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {AppMainComponent} from './app.main.component';
import {SharedModule} from "./shared/shared.module";
import {HttpClientModule} from "@angular/common/http";
import {ConfirmationService, MessageService} from "primeng/api";

@NgModule({
    imports: [
        BrowserModule,
        SharedModule,
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        AppRoutingModule,
        BrowserAnimationsModule,
    ],
    declarations: [
        AppComponent,
        AppMainComponent,
    ],
    providers: [
        {provide: LocationStrategy, useClass: HashLocationStrategy},
        MessageService, ConfirmationService, DatePipe,
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
