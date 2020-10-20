import {Injectable} from '@angular/core';
import {Observable, of} from 'rxjs';
import {ToastService} from "./toast.service";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {ArchivedNewsletter} from "./archived-newsletter";
import {NEWSLETTERS} from "./mock-archived-newsletters";
import {EnvService} from "./env.service";

@Injectable({
  providedIn: 'root'
})
export class ArchiveService {

  private listArchivedEntriesUrl = this.env.apiUrl + '/archivedentries';

  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'})
  }

  constructor(
    private http: HttpClient,
    private toastService: ToastService,
    private env: EnvService
  ) {
  }

  listArchivedEntries(): Observable<ArchivedNewsletter[]> {
    return of(NEWSLETTERS);
  }
}
