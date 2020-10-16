import {Component, ElementRef, OnInit, Renderer2, ViewChild} from '@angular/core';
import { ToastService } from "../toast.service";
import {Toast} from "../toast";
declare var $: any;

@Component({
  selector: 'app-toast',
  templateUrl: './toast.component.html',
  styleUrls: ['./toast.component.css']
})
export class ToastComponent implements OnInit {

  message: Toast;
  status: string;
  jqToast: any;
  private lastColoring: string;
  @ViewChild("header") header: ElementRef;
  @ViewChild("body") body: ElementRef;

  constructor(
    public toastService: ToastService,
    private renderer: Renderer2
  ) {
    this.message = new Toast("");
  }

  ngOnInit(): void {
    this.toastService.errorEvent.subscribe(msg => {
      this.message = msg;
      this.status = 'error';
      this.setColoring('danger');
      this.jqToast.toast('show');
    })
    this.toastService.successEvent.subscribe(msg => {
      this.message = msg;
      this.status = 'success';
      this.setColoring('success');
      this.jqToast.toast('show');
    })
    this.jqToast = $('.toast');
    this.jqToast.toast({
      delay: 2500
    });
    this.jqToast.on('hidden.bs.toast', () => {
      this.removeColoring();
    })
  }

  setColoring(coloring: string) {
    this.renderer.addClass(this.header.nativeElement, `border-${coloring}`);
    this.renderer.addClass(this.body.nativeElement, `border-${coloring}`);
    this.renderer.addClass(this.body.nativeElement, `bg-${coloring}`);
    this.lastColoring = coloring;
  }

  removeColoring() {
    this.renderer.removeClass(this.header.nativeElement, `border-${this.lastColoring}`);
    this.renderer.removeClass(this.body.nativeElement, `border-${this.lastColoring}`);
    this.renderer.removeClass(this.body.nativeElement, `bg-${this.lastColoring}`);
  }

}
