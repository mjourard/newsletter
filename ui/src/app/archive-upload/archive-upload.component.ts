import { Component, OnInit } from '@angular/core';
import {ArchiveService} from "../archive.service";
import {ArchivedNewsletter} from "../archived-newsletter";
import {FormsModule, NgForm, ReactiveFormsModule} from "@angular/forms";
import {OpResult} from "../op-result";
import {ToastService} from "../toast.service";

@Component({
  selector: 'app-archive-upload',
  templateUrl: './archive-upload.component.html',
  styleUrls: ['./archive-upload.component.css']
})
export class ArchiveUploadComponent implements OnInit {

  submitted = false;
  article: ArchivedNewsletter
  uploadResults: ArchivedNewsletter
  authors = [
    'John Adams',
    'Albert Einstein',
    'Thomas Jefferson'
  ]
  constructor(
    private archiveService: ArchiveService,
    private toastService: ToastService
  ) {
    this.reset(true);
  }

  get diagnostic() {
    let imgShort = this.article.img;
    imgShort = imgShort.substr(0, 50);
    if (this.article.img.length > 50) {
      imgShort += '...';
    }
    return JSON.stringify({
      abstract: this.article.abstract,
      author: this.article.author,
      title: this.article.title,
      img: imgShort
    });
  }

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
    this.archiveService.uploadArchivedEntry(this.article).subscribe(response => {
      if (response.hasOwnProperty('success')) {
        let opResult = response as OpResult
        if (opResult.success) {
          this.toastService.success(opResult.message)
        } else {
          this.toastService.error(opResult.message);
        }
      } else {
        this.uploadResults = response as ArchivedNewsletter
      }
    }, error => {
      this.toastService.error(error);
    });
  }

  reset(clear: boolean) {
    this.submitted = false;
    if (clear) {
      this.article = {
        abstract: "", author: "", img: "", title: ""
      }
      this.uploadResults = {
        abstract: "", author: "", img: "", title: ""
      }
    }
  }
}
