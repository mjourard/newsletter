import {Injectable} from '@angular/core';
import {Observable, of} from 'rxjs';
import {ToastService} from "./toast.service";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {ArchivedNewsletter} from "./archived-newsletter";
import {EnvService} from "./env.service";
import {LogService} from "./log.service";
import {OpResult} from "./op-result";
import {catchError, map} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class ArchiveService {

  private listArchivedEntriesUrl = this.env.apiUrl + '/archivedentries';
  private addArchivedEntryUrl = this.env.apiUrl + '/archivedarticle';

  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'})
  }

  private getArchivedEntriesUrl(max: number, lastId: string) {
    return this.listArchivedEntriesUrl + '?max=' + max + '&lastid=' + lastId
  }

  constructor(
    private http: HttpClient,
    private toastService: ToastService,
    private logger: LogService,
    private env: EnvService
  ) {
  }

  listArchivedEntries(): Observable<OpResult|ArchivedNewsletter[]> {
    return this.http.get<OpResult|ArchivedNewsletter[]>(this.getArchivedEntriesUrl(50, 'asdf'), this.httpOptions)
      .pipe(
        map(response => {
          return response as ArchivedNewsletter[];
        }),
        catchError(errResponse => {
          this.logger.error(errResponse);
          return of({success: false, message: errResponse.error.message} as OpResult);
        })
      );
  }

  uploadArchivedEntry(entry: ArchivedNewsletter): Observable<OpResult|ArchivedNewsletter> {
    console.log(this.addArchivedEntryUrl);
    return this.http.post<OpResult|ArchivedNewsletter>(this.addArchivedEntryUrl, entry, this.httpOptions)
      .pipe(
        map( response => {
          console.log(response);
          return response as ArchivedNewsletter;
        }),
        catchError(errResponse => {
          console.log(errResponse);
          this.logger.error(errResponse);
          return of({success: false, message: errResponse.error.message} as OpResult);
        })
      );
  }
}
