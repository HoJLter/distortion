import { Component, computed, effect, inject, Input } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';
import { ChannelModel, MessageModel } from '../../shared/models';
import { Message } from '../../shared/message/message';
import { ChatHeader } from './chat-header/chat-header';
import { ChatManager } from '../../shared/services/chat-manager';


@Component({
    imports: [ChatInput, Message, Message, ChatHeader],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {
    chatManager = inject(ChatManager);
    selectedChannel: ChannelModel | undefined;

    chatMessages = computed(() => {
        this.selectedChannel = this.chatManager.selectedTextChannel();
        return this.chatManager.getMessages(
            this.chatManager.selectedTextChannel()
        )
    })


}
