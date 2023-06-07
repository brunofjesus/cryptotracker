import {Injectable} from '@angular/core';
import {Subject} from 'rxjs';
import {MenuItem} from "primeng/api";

@Injectable({
    providedIn: 'root'
})
export class MenuService {

    private menuSource = new Subject<string>();
    private resetSource = new Subject();
    private menuItems = new Subject<MenuItem[]>();

    menuSource$ = this.menuSource.asObservable();
    resetSource$ = this.resetSource.asObservable();
    menuItems$ = this.menuItems.asObservable();

    onMenuStateChange(key: string) {
        this.menuSource.next(key);
    }

    setMenuItems(menuItems: MenuItem[]) {
        this.menuItems.next(menuItems);
    }

    reset() {
        //    this.resetSource.next(true);
    }
}
