import { Component, Input } from '@angular/core';

@Component({
  imports: [],
  selector: 'chat-channel',
  styleUrl: './chat-channel.css',
  templateUrl: './chat-channel.html',
})
export class ChatChannel {
    @Input() type: string = "";
    @Input() name: string = "";
}
