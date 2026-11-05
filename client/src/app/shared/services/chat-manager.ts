import { effect, Service, signal } from '@angular/core';
import { ChannelGroupModel, ChannelModel, MessageModel } from '../models';

@Service()
export class ChatManager {
    public selectedTextChannel = signal<ChannelModel | undefined>(undefined);
    public selectedVoiceChannel = signal<ChannelModel | undefined>(undefined);

    getMessages(channel: ChannelModel | undefined): MessageModel[] {
        if (!channel) {
            return [];
        }
        let messages: MessageModel[] = [];
        return messages;
    }

    sendMessage(channel: ChannelModel, msg: MessageModel) {}
}
