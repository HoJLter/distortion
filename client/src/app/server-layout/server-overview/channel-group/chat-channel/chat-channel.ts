import { Component, inject, Input, output } from '@angular/core';
import { ChatManager } from '../../../../shared/services/chat-manager';
import {ChannelModel} from '../../../../shared/data/models/channel.models';

@Component({
    imports: [],
    selector: 'chat-channel',
    styleUrl: './chat-channel.css',
    templateUrl: './chat-channel.html',
})
export class ChatChannel {
    @Input({ required: true }) channel!: ChannelModel;

    chatManager: ChatManager = inject(ChatManager);

    selectChannel(){
        if (this.channel.type == "text"){
            this.chatManager.selectedTextChannelUID.set(this.channel.uid);
        }
        else {
            this.chatManager.selectedVoiceChannelUID.set(this.channel.uid);
        }
    }
}
