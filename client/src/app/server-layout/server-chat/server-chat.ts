import { Component, computed, effect, inject, Input, signal } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';
import { ChatHeader } from './chat-header/chat-header';
import { ChatManager } from '../../shared/services/chat-manager';
import {MessageModel} from '../../shared/data/models/message.models';
import {Message} from '../../shared/message/message';


@Component({
    imports: [ChatInput, Message, ChatHeader],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {
    chatManager = inject(ChatManager);

    chatMessages = computed<MessageModel[]>((): MessageModel[] => {
        return this.chatManager.getMessages(this.chatManager.selectedTextChannelUID());
    });
}
