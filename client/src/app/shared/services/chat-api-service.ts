import { inject, Service } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GetMessageDTO, PostMessageDTO } from '../data/dto-models/message.models';

@Service()
export class ChatApiService {
    private httpClient = inject(HttpClient);

    // private serverIP: string = 'localhost:21530';
    private serverIP: string = '';

    getMessages(channelUID: string) {
        let messages: GetMessageDTO[] = [];

        this.httpClient
            .get(`${this.serverIP}/api/channels/${channelUID}/messages`)
            .subscribe((response) => {
                console.log(response);
            });

        return messages;
    }

    sendMessage(channelUID: string, message: PostMessageDTO) {
        console.log(message);
        console.log(JSON.stringify(message));

        this.httpClient
            .post(`${this.serverIP}/api/channels/${channelUID}/messages`, JSON.stringify(message))
            .subscribe((response) => {
                console.log(response);
            });
    }
}
