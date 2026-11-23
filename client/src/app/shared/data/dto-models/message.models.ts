

export interface MessageAuthorDTO {
    uid: string;
    username: string;
    avatar_path: string;
}

export interface GetMessageDTO {
    message_id: number;
    author: MessageAuthorDTO;
    content: string;
}

export interface PostMessageDTO {
    channel_uid: string;
    user_uid: string;
    content: string;
}

export interface WSMessage {
    type: string;
    payload: Object;
}
