import { GetMessageDTO } from './dto-models/message.models';
import { MessageModel } from './models/message.models';


export function mapMessage(dtoMessage: GetMessageDTO): MessageModel{
    let mappedMessage: MessageModel = {
        id: dtoMessage.message_id,
        author: {
            username: dtoMessage.author.username,
            avatarPath: dtoMessage.author.avatar_path,
        },
        content: dtoMessage.content,
    };

    console.log("mapping");
    return mappedMessage
}
