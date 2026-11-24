import { effect, inject, Service, signal } from '@angular/core';
import { MessageModel } from '../../data/models/message.models';
import { DbApiService } from '../transport/db-api-service';
import { GetMessageResponse, PostMessageRequest } from '../../data/dto-models/message.models';
import { mapChannel, mapMessage } from '../../data/mappers';
import { filter, finalize, map, Observable, tap } from 'rxjs';
import { UserService } from './user-service';
import { WebsocketService, WSMessageType } from '../transport/websocket-service';
import { HttpClient } from '@angular/common/http';
import { ChannelModel } from '../../data/models/channel.models';


@Service()
export class ChatManager {
    messages = signal<MessageModel[]>([]);
    isLoading = signal<boolean>(false);

    selectedTextChannel = signal<string>("choose the channel")
    channels = signal<ChannelModel[]>([]);

    userService: UserService = inject(UserService);
    websocketService = inject(WebsocketService);
    dbApi: DbApiService = inject(DbApiService);

    constructor(){
        this.websocketService
            .observe()
            .pipe(
                tap((msg) => console.log('INITIAL', msg)),
                filter((wsMsg) => wsMsg.type === WSMessageType.broadcastMessage),
                tap((msg) => console.log('AFTER FILTER', msg)),
                map((wsMsg) => wsMsg.payload as GetMessageResponse),
                tap((msg) => console.log('AFTER CONVERSION', msg)),
                map((getMsgDTO) => mapMessage(getMsgDTO)),
                tap((msg) => console.log('AFTER MAPPING', msg)),
            )
            .subscribe((message) => {
                this.messages.update((msgs) => [...msgs, message]);
                console.log('Message acquired');
            });
    }

    updateChannels() {
        console.log("update channels called");
        this.dbApi.getChannels().pipe(
            map((channels) => channels.map(mapChannel)),
            map((channels) => channels.filter((channel) => channel.type == "text")),
        ).subscribe((channels => {
            console.log(channels)
            this.channels.set(channels);
        }))
    }

    selectTextChannel(channelUID: string){
        this.selectedTextChannel.set(channelUID);
        this.messages.set([]);
        this.loadMessages(this.selectedTextChannel(), 0, 50);
    }

    loadMessages(channelUID: string, offset: number, limit: number) {
        if (this.isLoading()) {
            return;
        }
        this.isLoading.set(true);

        this.dbApi.getMessages(channelUID, offset, limit)
            .pipe(
                map((dto => dto.map(mapMessage))),
                finalize(() => this.isLoading.set(false))
            )
            .subscribe((newMessages: MessageModel[]) => {
            this.messages.update((currentMessages) => [...newMessages, ...currentMessages] );
        });
    }

    // необходимо вебсокетом ловить и это получение сообщения
    sendMessage(channelUID: string, message: string): void {
        let dtoMessage: PostMessageRequest = {
            channel_uid: channelUID,
            user_uid: this.userService.getUID(),
            content: message,
        };
        console.log(dtoMessage);

        this.websocketService.sendMessage(dtoMessage);
    }
}
