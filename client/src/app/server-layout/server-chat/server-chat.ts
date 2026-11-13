import { Component, computed, effect, inject, Input, signal } from '@angular/core';
import { ChatInput } from './chat-input/chat-input';
import { ChatHeader } from './chat-header/chat-header';
import { ChatManager } from '../../shared/services/chat-manager';
import {MessageModel} from '../../shared/data/models/message.models';
import {Message} from '../../shared/message/message';
import { ChatMessages } from './chat-messages/chat-messages';


@Component({
    imports: [ChatInput, ChatHeader, ChatMessages],
    selector: 'server-chat',
    styleUrl: './server-chat.css',
    templateUrl: './server-chat.html',
})
export class ServerChat {}
