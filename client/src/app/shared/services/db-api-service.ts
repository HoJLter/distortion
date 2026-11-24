import { inject, Service } from '@angular/core';
import { HttpClient, HttpResponse} from '@angular/common/http';
import {
    GetMessageResponse,
    PostMessageRequest,
} from '../data/dto-models/message.models';
import { Observable } from 'rxjs';
import {CreateChannelRequest} from "../data/dto-models/channel.models";
import {GetUserResponse} from "../data/dto-models/user.models";

@Service()
export class DbApiService {
    private httpClient = inject(HttpClient);

    // private serverIP: string = 'localhost:21530';
    private serverIP: string = '';

    getMessages(
        channelUID: string,
        offset: number,
        limit: number,
    ): Observable<GetMessageResponse[]> {
        console.log('getMessages called');
        return this.httpClient.get<GetMessageResponse[]>(
            `${this.serverIP}/api/channels/${channelUID}/messages?offset=${offset}&limit=${limit}`,
        );
    }

    createChannel(channelUID: string, type: string): Observable<void> {
        console.log('createChannel called');
        let body: CreateChannelRequest = {
            channel_uid: channelUID,
            type: type,
        };
        return this.httpClient.post<void>(`${this.serverIP}/api/channels`, body);
    }

    getChannels(): Observable<GetMessageResponse[]> {
        console.log('getChannels called');
        return this.httpClient.get<GetMessageResponse[]>(`${this.serverIP}/api/channels`);
    }

    getUsers(): Observable<GetUserResponse[]> {
        console.log('getUsers');
        return this.httpClient.get<GetUserResponse[]>(`${this.serverIP}/api/users`);
    }
}
