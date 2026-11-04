
export interface MemberModel {
    username: string;
}

export interface UserModel {
    name: string;
}

export interface MessageModel {
    user: UserModel;
    msg: string;
}

export interface ServerNavModel {
    iconPath: string;
}

export interface ChannelModel {
    name: string;
    type: string;
    users: MemberModel[];
}

export interface ChannelGroupModel {
    name: string;
    type: string;
    channels: ChannelModel[];
}

export interface MemberGroupModel {
    name: string;
    members: MemberModel[];
}
