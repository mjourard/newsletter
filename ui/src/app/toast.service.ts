import { Injectable } from '@angular/core';
import {Toast} from "./toast";

@Injectable({
  providedIn: 'root'
})
export class ToastService {

  messages: Toast[] = [];

  add(message: string) {
    let toast = new Toast(message);
    this.messages.push(toast);
  }

  clear(id: string) {
    this.messages = this.messages.filter(toast => toast.id !== id);
  }
  constructor() { }
}
