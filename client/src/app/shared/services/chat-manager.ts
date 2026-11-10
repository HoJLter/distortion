import { effect, inject, Service, signal } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import { GetMessageDTO, PostMessageDTO } from '../data/dto-models/message.models';
import { MessageModel } from '../data/models/message.models';
import { ChatApiService } from './chat-api-service';
import { mapMessage } from '../data/mappers';
import { UserService } from './user-service';

@Service()
export class ChatManager {
    public selectedTextChannelUID = signal<string>('');
    public selectedVoiceChannelUID = signal<string>('');

    private chatApiService = inject(ChatApiService);
    private userService = inject(UserService);

    getMessages(channelUID: string): MessageModel[] {
        let dtoMessages: GetMessageDTO[] = this.chatApiService.getMessages(channelUID);
        return dtoMessages.map((dtoMsg) => mapMessage(dtoMsg));
    }

    sendMessage(channelUID: string, message: string): void {
        let dtoMessage: PostMessageDTO = {
            user_uid: this.userService.getUID(),
            content: message,
        };
        this.chatApiService.sendMessage(channelUID, dtoMessage);
    }
}
