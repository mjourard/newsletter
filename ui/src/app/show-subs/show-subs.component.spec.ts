import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowSubsComponent } from './show-subs.component';

describe('ShowSubsComponent', () => {
  let component: ShowSubsComponent;
  let fixture: ComponentFixture<ShowSubsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ShowSubsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ShowSubsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
