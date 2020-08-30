import { Injectable } from '@angular/core';
import {Subscriber} from "./subscriber";
import { SUBSCRIBERS } from "./mock-subscribers";
import { Observable, of } from 'rxjs';
import { ToastService } from "./toast.service";
import { HttpClient, HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class SubscriptionService {

  private baseUrl = 'https://du7hl2x8w2.execute-api.us-east-1.amazonaws.com/dev';
  private listSubsUrl = this.baseUrl + '/getemails';
  private subscribeEmailUrl = this.baseUrl + '/email';

  constructor(
    private http: HttpClient,
    private toastService: ToastService
  ) { }

  subscribeEmail(): void {

  }

  listSubscribers(): Observable<Subscriber[]> {
    // TODO: send the message _after_ fetching the heroes
    // this.log('subscriptionService: fetched subs');
    // return of(SUBSCRIBERS);
    return this.http.get<Subscriber[]>(this.listSubsUrl);
  }

  private log(message: string) {
    this.toastService.add(`SubscriptionService: ${message}`);
  }
}
