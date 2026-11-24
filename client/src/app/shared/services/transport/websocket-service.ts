import { Service, Signal, WritableSignal } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import { Message } from '../../message/message';
import { PostMessageRequest, WSMessage } from '../../data/dto-models/message.models';
import {MessageModel} from "../../data/models/message.models";
import {webSocket} from "rxjs/webSocket";

export namespace WSMessageType {
    export const broadcastMessage = "broadcastMessage";
    export const postMessage = "postMessage";
}

@Service()
export class WebsocketService {
    private socket = webSocket<WSMessage>("ws://127.0.0.1:21530/api/ws");

    sendMessage(message: PostMessageRequest) {
        let wsMessage = {
            type: WSMessageType.postMessage,
            payload: message
        }
        this.socket.next(wsMessage);
    }

    observe() {
        return this.socket.asObservable()
    }
}
