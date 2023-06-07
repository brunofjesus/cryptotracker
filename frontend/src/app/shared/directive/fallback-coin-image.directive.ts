import {Directive, ElementRef, HostListener, Input} from '@angular/core';

@Directive({
    selector: 'img[appFallbackCoinImage]'
})
export class FallbackCoinImageDirective {

    @Input() appImgFallback: string;

    constructor(
        private eRef: ElementRef
    ) {
    }

    @HostListener('error')
    loadFallbackOnError() {
        const element: HTMLImageElement = <HTMLImageElement>this.eRef.nativeElement;
        element.src = this.appImgFallback || 'assets/currency/generic.svg';
    }

}
