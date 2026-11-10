

export interface MessageAuthorModel {
    username: string;
    avatarPath: string;
}

export interface MessageModel{
    author: MessageAuthorModel;
    content: string;
}
