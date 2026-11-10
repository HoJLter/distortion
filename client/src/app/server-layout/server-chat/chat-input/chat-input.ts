import { Component, EventEmitter, inject, Output } from '@angular/core';
import { ChatManager } from '../../../shared/services/chat-manager';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { UserService } from '../../../shared/services/user-service';
import {PostMessageDTO} from '../../../shared/data/dto-models/message.models';

@Component({
    imports: [ReactiveFormsModule],
    selector: 'chat-input',
    styleUrl: './chat-input.css',
    templateUrl: './chat-input.html',
})
export class ChatInput {
    @Output() messageChange = new EventEmitter();
    chatManager: ChatManager = inject(ChatManager);
    userService = inject(UserService);

    chatInput: FormControl = new FormControl();

    sendMessage() {
        let channelUID = this.chatManager.selectedTextChannelUID();
        if (channelUID == '') {
            alert('Choose the channel to send the message!');
            return;
        }

        this.chatManager.sendMessage(channelUID, this.chatInput.value);
        this.messageChange.emit();
        this.chatInput.reset();
    }
}
