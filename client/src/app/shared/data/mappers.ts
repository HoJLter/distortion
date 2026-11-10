import { GetMessageDTO, MessageAuthorDTO } from './dto-models/message.models';
import { MessageModel } from './models/message.models';


export function mapMessage(dtoMessage: GetMessageDTO): MessageModel{
    return {
        author: {
            username: dtoMessage.author.username,
            avatarPath: dtoMessage.author.avatar_path,
        },
        content: dtoMessage.content,
    };
}
