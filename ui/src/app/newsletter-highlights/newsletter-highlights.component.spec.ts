import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewsletterHighlightsComponent } from './newsletter-highlights.component';
import {ArchivedNewsletter} from "../archived-newsletter";
import {ArchiveService} from "../archive.service";
import {of} from "rxjs";

describe('NewsletterHighlightsComponent', () => {
  let component: NewsletterHighlightsComponent;
  let fixture: ComponentFixture<NewsletterHighlightsComponent>;
  let archivedEntries: ArchivedNewsletter[] = [
    {
      title: 'Test Archive',
      img: 'https://someimage.com/YYXXZZ',
      abstract: 'A test archive entry',
      author: 'Einstein',
    }
  ];

  beforeEach(async(() => {

    const archiveService = jasmine.createSpyObj('ArchiveService', ['listArchivedEntries']);
    archiveService.listArchivedEntries.and.returnValue(of(archivedEntries));
    TestBed.configureTestingModule({
      declarations: [ NewsletterHighlightsComponent ],
      providers: [{provide: ArchiveService, useValue: archiveService}]
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
