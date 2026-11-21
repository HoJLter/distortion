import { Service, Signal, WritableSignal } from '@angular/core';
import { webSocket } from 'rxjs/webSocket';
import { GetMessageDTO, PostMessageDTO } from '../data/dto-models/message.models';
import { Subscription } from 'rxjs';
import { MessageModel } from '../data/models/message.models';
import { mapMessage } from '../data/mappers';

@Service()
export class WebsocketService {
    private socket = webSocket<unknown>({ url: 'ws://localhost:21530/api/ws'});

    subscribe(messages: WritableSignal<MessageModel[]>) {

        return this.socket.subscribe((msg) => {
            messages.update((msgs) => [...msgs, mapMessage(msg as GetMessageDTO)]);
        });
    }

    sendMessage(message: PostMessageDTO) {
        this.socket.next(message);
    }
}
