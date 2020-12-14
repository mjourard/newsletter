import { NavigationEnd } from '@angular/router';
import {Observable} from "rxjs";

export class MockRouter {
  public ne = new NavigationEnd(0, 'http://localhost:4200/login', 'http://localhost:4200/login');
  public events = new Observable(observer => {
    observer.next(this.ne);
    observer.complete();
  });
}

class MockRouterNoLogin {
  public ne = new NavigationEnd(0, 'http://localhost:4200/dashboard', 'http://localhost:4200/dashboard');
  public events = new Observable(observer => {
    observer.next(this.ne);
    observer.complete();
  });
}
