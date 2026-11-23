import { inject, Service } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GetMessageResponse, PostMessageRequest } from '../data/dto-models/message.models';
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

    getMessages(channelUID: string, offset: number, limit: number): Observable<GetMessageResponse[]> {
        console.log('getMessages called');
        return this.httpClient.get<GetMessageResponse[]>(
            `${this.serverIP}/api/channels/${channelUID}/messages?offset=${offset}&limit=${limit}`,
        );
    }

    createChannel(channelUID: string, type: string): Observable<GetMessageResponse[]> {
        console.log('getUsers');
        return this.httpClient.get<GetMessageResponse[]>(
            `${this.serverIP}/api/channels/`,
        );
    }

    getChannels(): Observable<GetMessageResponse[]> {
        console.log('getUsers');
        return this.httpClient.get<GetMessageResponse[]>(
            `${this.serverIP}/api/channels/`,
        );
    }

    getUsers(): Observable<GetMessageResponse[]> {
        console.log('getUsers');
        return this.httpClient.get<GetMessageResponse[]>(
            `${this.serverIP}/api/users`,
        );
    }
}
