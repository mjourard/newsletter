import { Injectable } from '@angular/core';
import {Subscriber} from "./subscriber";
import { Observable, of } from 'rxjs';
import { ToastService } from "./toast.service";
import { HttpClient, HttpHeaders} from "@angular/common/http";
import { catchError, map, tap } from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class SubscriptionService {

  private baseUrl = 'https://du7hl2x8w2.execute-api.us-east-1.amazonaws.com/dev';
  private listSubsUrl = this.baseUrl + '/getemails';
  private subscribeEmailUrl = this.baseUrl + '/email';

  private httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json'})
  }

  constructor(
    private http: HttpClient,
    private toastService: ToastService
  ) { }

  subscribeEmail(email: string): Observable<any> {
    return this.http.post(this.subscribeEmailUrl, {email: email}, this.httpOptions).pipe(
      tap(_ => this.log(`subscribing ${email}`)),
      catchError(this.handleError<any>('subscribeEmail'))
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
    this.toastService.add(`SubscriptionService: ${message}`);
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (response: any): Observable<T> => {
      // TODO: send the error to remote logging infrastructure
      console.error(response); // log to console instead
      this.log(response.error.message);
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
