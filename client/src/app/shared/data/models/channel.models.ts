import { UserModel } from './user.models';

export interface ChannelModel {
    uid: string;
    type: string;
    users: UserModel[];
}

export interface ChannelGroupModel {
    name: string;
    type: string;
    channels: ChannelModel[];
}
