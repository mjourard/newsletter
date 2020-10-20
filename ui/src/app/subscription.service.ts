import {Injectable} from '@angular/core';
import {Subscriber} from "./subscriber";
import {Observable, of} from 'rxjs';
import {ToastService} from "./toast.service";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {catchError, map, tap} from "rxjs/operators";
import {LogService} from "./log.service";
import {OpResult} from "./op-result";
import {EnvService} from "./env.service";

@Injectable({
  providedIn: 'root'
})
export class SubscriptionService {

  private listSubsUrl = this.env.apiUrl + '/getemails';
  private subscribeEmailUrl = this.env.apiUrl + '/email';

  private httpOptions = {
    headers: new HttpHeaders({'Content-Type': 'application/json'})
  }

  constructor(
    private http: HttpClient,
    private toastService: ToastService,
    private logger: LogService,
    private env: EnvService
  ) {
  }

  subscribeEmail(email: string): Observable<OpResult> {
    return this.http.post<OpResult>(this.subscribeEmailUrl, {email: email}, this.httpOptions)
      .pipe(
        map( _ => {
          return {success: true, message: `Successfully subscribed ${email}`} as OpResult;
        }),
        catchError(errResponse => {
          this.logger.error(errResponse);
          return of({success: false, message: errResponse.error.message} as OpResult);
        })
    );
  }

  listSubscribers(): Observable<Subscriber[]> {
    return this.http.get<Subscriber[]>(this.listSubsUrl)
      .pipe(
        tap(_ => this.log("fetched subscribers")),
        catchError(this.handleError<Subscriber[]>('listSubscribers', []))
      );
  }

  private log(message: string) {
    this.toastService.success(`SubscriptionService: ${message}`);
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (response: any): Observable<T> => {
      this.logger.error(response); // log to console instead
      this.log(response.error.message);
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
