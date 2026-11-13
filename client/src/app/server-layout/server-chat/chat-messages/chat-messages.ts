import { Component, effect, inject, untracked } from '@angular/core';
import { ChatManager } from '../../../shared/services/chat-manager';
import { Message } from '../../../shared/message/message';

@Component({
    imports: [Message],
    selector: 'chat-messages',
    styleUrl: './chat-messages.css',
    templateUrl: './chat-messages.html',
})
export class ChatMessages {
    chatManager = inject(ChatManager);
    chatMessages = this.chatManager.messages;

    // constructor() {
    //     effect(() => {
    //         console.log('Channel switched');
    //         untracked(() => {
    //             this.chatManager.loadMessages(this.chatManager.selectedTextChannel(), 0, 50);
    //         })
    //     })
    // }

}
