import {Injectable} from '@angular/core';
import {Subject} from "rxjs";
import {Event} from "./model/event";

@Injectable({
    providedIn: 'root'
})
export class EventService {

    private eventSubject = new Subject<Event>()

    event$ = this.eventSubject.asObservable();

    emitEvent(event: Event) {
        this.eventSubject.next(event);
    }

    constructor() {
    }
}
