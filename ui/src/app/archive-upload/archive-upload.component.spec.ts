import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from "@angular/forms";
import { ArchiveUploadComponent } from './archive-upload.component';
import {ArchiveService} from "../archive.service";

describe('ArchiveUploadComponent', () => {
  let component: ArchiveUploadComponent;
  let fixture: ComponentFixture<ArchiveUploadComponent>;

  beforeEach(async(() => {
    const archiveService = jasmine.createSpyObj('ArchiveService', ['uploadArchivedEntry']);
    TestBed.configureTestingModule({
      imports: [FormsModule],
      declarations: [ ArchiveUploadComponent ],
      providers: [{provide: ArchiveService, useValue: archiveService}]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ArchiveUploadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
