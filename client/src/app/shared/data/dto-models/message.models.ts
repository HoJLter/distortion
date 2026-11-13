

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
    user_uid: string,
    content: string,
}
