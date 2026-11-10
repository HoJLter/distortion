import { effect, inject, Service, signal } from '@angular/core';
import { ChannelGroupModel, ChannelModel, MessageViewModel } from '../models';
import {HttpClient} from "@angular/common/http";
import {PostMessageDTO} from "../dto-models";

@Service()
export class ChatManager {
    public selectedTextChannelUID = signal<string>('');
    public selectedVoiceChannelUID = signal<string>('');

    // private serverIP: string = 'localhost:21530';
    private serverIP: string = '';
    private httpClient = inject(HttpClient);

    getMessages(channelUID: string): MessageViewModel[] {
        let messages: MessageViewModel[] = [];
        this.httpClient
            .get(`${this.serverIP}/api/channels/${channelUID}/messages`)
            .subscribe((response) => {
                console.log(response);
            });

        return messages;
    }

    sendMessage(channelUID: string, message: PostMessageDTO) {
        console.log(message)
        console.log(JSON.stringify(message));
        this.C.update((list) => [...list, created]);

        this.httpClient
            .post(`${this.serverIP}/api/channels/${channelUID}/messages`, JSON.stringify(message))
            .subscribe((response) => {
                console.log(response);
            });
    }
}
