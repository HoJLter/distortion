import { MessageAuthorDTO } from './message.models';

export interface GetChannelResponse {
    message_id: number;
    author: MessageAuthorDTO;
    content: string;
}

export interface CreateChannelRequest {
    channel_uid: string;
    type: string;
}
