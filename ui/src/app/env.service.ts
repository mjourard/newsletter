import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class EnvService {

  public apiUrl = '';

  // Whether or not to enable debug mode
  // Setting this to false will disable console output
  public enableDebug = true;
  public callToSubscribeDevelopment = true;
  constructor() { }
}
