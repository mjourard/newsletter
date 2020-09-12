import {Injectable} from '@angular/core';
import {Observable, of} from 'rxjs';
import {ToastService} from "./toast.service";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {ArchivedNewsletter} from "./archived-newsletter";
import {NEWSLETTERS} from "./mock-archived-newsletters";

@Injectable({
  providedIn: 'root'
})
export class ArchiveService {

  private baseUrl = 'https://du7hl2x8w2.execute-api.us-east-1.amazonaws.com/dev';
  private listArchivedEntriesUrl = this.baseUrl + '/archivedentries';

  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'})
  }

  constructor(
    private http: HttpClient,
    private toastService: ToastService
  ) {
  }

  listArchivedEntries(): Observable<ArchivedNewsletter[]> {
    return of(NEWSLETTERS);
  }
}
