import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewsletterHighlightsComponent } from './newsletter-highlights.component';

describe('NewsletterHighlightsComponent', () => {
  let component: NewsletterHighlightsComponent;
  let fixture: ComponentFixture<NewsletterHighlightsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewsletterHighlightsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewsletterHighlightsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
