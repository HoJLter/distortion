
export interface MemberModel {
    username: string;
}

export interface UserViewModel {
    name: string;
}

export interface MessageViewModel {
    user: UserViewModel;
    content: string;
}

export interface ServerNavModel {
    iconPath: string;
}

export interface ChannelModel {
    uid: string;
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
