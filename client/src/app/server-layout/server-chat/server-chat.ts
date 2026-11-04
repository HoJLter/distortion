import { Component } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';
import { ChannelModel, MessageModel } from '../../shared/models';
import { Message } from '../../shared/message/message';
import { ChatHeader } from './chat-header/chat-header';


@Component({
    imports: [ChatInput, Message, Message, ChatHeader],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {
    chatMessages: MessageModel[] = [];
    currentChannel: ChannelModel | undefined;
}
