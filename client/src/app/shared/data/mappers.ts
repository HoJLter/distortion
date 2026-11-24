import { GetMessageResponse } from './dto-models/message.models';
import { MessageModel } from './models/message.models';
import { GetUserResponse } from './dto-models/user.models';
import { UserModel } from './models/user.models';
import {GetChannelResponse} from "./dto-models/channel.models";
import {ChannelModel} from "./models/channel.models";


export function mapMessage(dtoMessage: GetMessageResponse): MessageModel{
    let mappedMessage: MessageModel = {
        id: dtoMessage.message_id,
        author: {
            username: dtoMessage.author.username,
            avatarPath: dtoMessage.author.avatar_path,
        },
        content: dtoMessage.content,
    };

    console.log("message mapping called");
    return mappedMessage
}

export function mapUser(dtoUser: GetUserResponse): UserModel {
    let user: UserModel = {
        username: dtoUser.username
    }
    console.log('user mapping called');
    return user
}

export function mapChannel(dtoChannel: GetChannelResponse): ChannelModel {
    let channel: ChannelModel = {
        uid: dtoChannel.channel_uid,
        type: dtoChannel.type,
    }
    console.log('channel mapping called');
    return channel;
}
