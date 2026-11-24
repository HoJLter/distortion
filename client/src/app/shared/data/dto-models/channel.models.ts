import { MessageAuthorDTO } from './message.models';

export interface GetChannelResponse {
    channel_uid: string;
    type: string;
}

export interface CreateChannelRequest {
    channel_uid: string;
    type: string;
}
