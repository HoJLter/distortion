

export interface MessageAuthorDTO {
    uid: string;
    username: string;
    avatar_path: string;
}

export interface GetMessageResponse {
    message_id: number;
    author: MessageAuthorDTO;
    content: string;
}

export interface PostMessageRequest {
    channel_uid: string;
    user_uid: string;
    content: string;
}

export interface WSMessage {
    type: string;
    payload: Object;
}

export interface GetChannelResponse {
    message_id: number;
    author: MessageAuthorDTO;
    content: string;
}

export interface CreateChannelRequest {
    channel_uid: string;
    type: string;
}

export interface GetUserResponse {
    message_id: number;
    author: MessageAuthorDTO;
    content: string;
}


