import { Component, inject } from '@angular/core';
import { ChatManager } from '../../../shared/services/presentation/chat-manager';
import { ChatChannel } from './chat-channel/chat-channel';

@Component({
    imports: [ChatChannel],
    selector: 'channels-block',
    styleUrl: './channels-block.css',
    templateUrl: './channels-block.html',
})
export class ChannelsBlock {
    chatManager = inject(ChatManager);

    constructor() {
        this.chatManager.updateChannels()
    }

}
