export class Toast {
  message: string;
  timestamp: Date;
  id: string;

  constructor(message: string) {
    this.message = message;
    this.timestamp = new Date();
    this.id = window.crypto.getRandomValues(new Uint32Array(1))[0].toString(16);
  }

  getUserFriendlyDT() {
    return [
      this.timestamp.getFullYear(),
      this.timestamp.getMonth(),
      this.timestamp.getDate()
      ].join("-") + " " +
      [
        this.timestamp.getHours(),
        this.timestamp.getMinutes(),
        this.timestamp.getSeconds()
      ].join(":")
  }
}
