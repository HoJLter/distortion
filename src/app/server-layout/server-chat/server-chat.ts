import { Component } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';

@Component({
    imports: [ChatInput],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {}
