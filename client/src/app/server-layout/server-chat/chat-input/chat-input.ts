import { Component, inject } from '@angular/core';
import { ChatManager } from '../../../shared/services/chat-manager';

@Component({
    imports: [],
    selector: 'chat-input',
    styleUrl: './chat-input.css',
    templateUrl: './chat-input.html',
})
export class ChatInput {
    chatManager: ChatManager = inject(ChatManager);
}
