import { Service, Signal, WritableSignal } from '@angular/core';
import { PostMessageRequest, WSMessage } from '../../data/dto-models/message.models';
import {webSocket} from "rxjs/webSocket";

export namespace WSMessageType {
    export const broadcastMessage = "broadcastMessage";
    export const postMessage = "postMessage";
}

@Service()
export class WebsocketService {
    private socket = webSocket<WSMessage>("");

    changeChannel(channelId: string) {
        this.socket.complete();
        this.socket = webSocket<WSMessage>(`ws://127.0.0.1:21530/api/ws/${channelId}`);
    }

    sendMessage(message: PostMessageRequest) {
        let wsMessage = {
            type: WSMessageType.postMessage,
            payload: message,
        };
        this.socket.next(wsMessage);
    }

    observe() {
        return this.socket.asObservable();
    }
}
