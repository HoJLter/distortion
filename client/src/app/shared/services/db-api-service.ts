import { inject, Service } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GetMessageDTO, PostMessageDTO } from '../data/dto-models/message.models';
import { Observable } from 'rxjs';

@Service()
export class DbApiService {
    private httpClient = inject(HttpClient);

    // private serverIP: string = 'localhost:21530';
    private serverIP: string = '';

    // getMessages(channelUID: string): Observable<GetMessageDTO[]> {
    //     return this.httpClient.get<GetMessageDTO[]>(
    //         `${this.serverIP}/api/channels/${channelUID}/messages`,
    //     );
    // }

    getMessages(channelUID: string, offset: number, limit: number): Observable<GetMessageDTO[]> {
        console.log('getMessages');
        return this.httpClient.get<GetMessageDTO[]>(
            `${this.serverIP}/api/channels/${channelUID}/messages?offset=${offset}&limit=${limit}`,
        );
    }

    sendMessage(channelUID: string, message: PostMessageDTO) {
        this.httpClient
            .post(`${this.serverIP}/api/channels/${channelUID}/messages`, JSON.stringify(message))
            .subscribe((response) => {
                console.log(response);
            });
    }
}
