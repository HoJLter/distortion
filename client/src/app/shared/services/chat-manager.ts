import { effect, inject, Service, signal } from '@angular/core';
import { MessageModel } from '../data/models/message.models';
import { DbApiService } from './db-api-service';
import { GetMessageDTO, PostMessageDTO } from '../data/dto-models/message.models';
import { mapMessage } from '../data/mappers';
import { finalize, map } from 'rxjs';
import { UserService } from './user-service';


@Service()
export class ChatManager {
    messages = signal<MessageModel[]>([]);
    isLoading = signal<boolean>(false);

    selectedTextChannel = signal<string>("undefined")

    userService: UserService = inject(UserService);
    dbApi: DbApiService = inject(DbApiService);

    selectTextChannel(channelUID: string){
        this.selectedTextChannel.set(channelUID);
        this.messages.set([]);
        this.loadMessages(this.selectedTextChannel(), 0, 50);
    }

    loadMessages(channelUID: string, offset: number, limit: number) {
        if (this.isLoading()) {
            return;
        }
        this.isLoading.set(true);

        this.dbApi.getMessages(channelUID, offset, limit)
            .pipe(
                map((dto => dto.map(mapMessage))),
                finalize(() => this.isLoading.set(false))
            )
            .subscribe((newMessages: MessageModel[]) => {
            this.messages.update((currentMessages) => [...newMessages, ...currentMessages] );
        });
    }

    // необходимо вебсокетом ловить и это получение сообщения
    sendMessage(channelUID: string, message: string): void {
        let dtoMessage: PostMessageDTO = {
            user_uid: this.userService.getUID(),
            content: message,
        };

        this.dbApi.sendMessage(channelUID, dtoMessage)
    }
}
