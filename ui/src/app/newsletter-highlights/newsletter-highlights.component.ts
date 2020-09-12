import { Component, OnInit } from '@angular/core';
import {ArchivedNewsletter} from "../archived-newsletter";
import {ArchiveService} from "../archive.service";

@Component({
  selector: 'app-newsletter-highlights',
  templateUrl: './newsletter-highlights.component.html',
  styleUrls: ['./newsletter-highlights.component.css']
})
export class NewsletterHighlightsComponent implements OnInit {

  newsletters: ArchivedNewsletter[];

  constructor(
    private archiveService: ArchiveService
  ) {
  }

  ngOnInit(): void {
    this.listArchivedEntries();
  }

  listArchivedEntries(): void {
    this.archiveService.listArchivedEntries().subscribe(
      entries => this.newsletters = entries
    );
  }
}
