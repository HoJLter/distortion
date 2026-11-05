import { Component, inject, Input, output } from '@angular/core';
import { ChannelModel } from '../../../../shared/models';
import { ChatManager } from '../../../../shared/services/chat-manager';

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
            this.chatManager.selectedTextChannel.set(this.channel);
        }
        else {
            this.chatManager.selectedVoiceChannel.set(this.channel);
        }
    }
}
