

export interface MessageAuthorModel {
    username: string;
    avatarPath: string;
}

export interface MessageModel{
    id: number;
    author: MessageAuthorModel;
    content: string;
}
