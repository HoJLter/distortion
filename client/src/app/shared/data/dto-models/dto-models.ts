
export interface MessageAuthorDTO {
    username: string;
    avatarPath: string;
}


export interface MessageDTO {
    id: string;
    author: MessageAuthorDTO;
    content: string;
}

export interface PostMessageDTO {
    user_uid: string;
    content: string;
}
