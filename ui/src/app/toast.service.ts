import {Injectable, Output, EventEmitter} from '@angular/core';
import {Toast} from "./toast";

@Injectable({
  providedIn: 'root'
})
export class ToastService {
  @Output() successEvent = new EventEmitter<Toast>();
  @Output() errorEvent = new EventEmitter<Toast>();
  success(message: string) {
    this.successEvent.emit(new Toast(message));
  }
  error(message: string) {
    this.errorEvent.emit(new Toast(message));
  }

  constructor() {
  }
}
