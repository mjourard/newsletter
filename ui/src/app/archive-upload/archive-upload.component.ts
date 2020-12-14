import { Component, OnInit } from '@angular/core';
import {ArchiveService} from "../archive.service";
import {ArchivedNewsletter} from "../archived-newsletter";
import {FormsModule, NgForm, ReactiveFormsModule} from "@angular/forms";

@Component({
  selector: 'app-archive-upload',
  templateUrl: './archive-upload.component.html',
  styleUrls: ['./archive-upload.component.css']
})
export class ArchiveUploadComponent implements OnInit {

  submitted = false;
  article: ArchivedNewsletter
  authors = [
    'John Adams',
    'Albert Einstein',
    'Thomas Jefferson'
  ]
  constructor(
    private archiveService: ArchiveService
  ) {
    this.article = {
      abstract: "", author: "", img: "", title: ""
    }
  }

  get diagnostic() { return JSON.stringify(this.article); }

  ngOnInit(): void {
  }

  fileSelected(fileList) {
    let reader = new FileReader();
    reader.onload = (e) => {
      if (typeof e.target.result === "string") {
        this.article.img = e.target.result;
      }
    }
    reader.readAsDataURL(fileList[0]);
  }

  onSubmit(archivedUploadForm: NgForm) {
    //TODO: properly handle image uploading such that it populates the preview space. This should deal with it:
    //https://academind.com/learn/angular/snippets/angular-image-upload-made-easy/
    this.submitted = true;
    console.log(this.article);
    this.archiveService.uploadArchivedEntry(this.article)
  }
}
