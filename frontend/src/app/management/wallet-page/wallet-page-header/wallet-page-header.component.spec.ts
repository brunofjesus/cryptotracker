import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WalletPageHeaderComponent } from './wallet-page-header.component';

describe('WalletPageHeaderComponent', () => {
  let component: WalletPageHeaderComponent;
  let fixture: ComponentFixture<WalletPageHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WalletPageHeaderComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WalletPageHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
