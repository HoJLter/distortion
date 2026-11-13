import { Component, Input } from '@angular/core';
import { MessageModel } from '../data/models/message.models';

@Component({
    imports: [],
    selector: 'message',
    styleUrl: './message.css',
    templateUrl: './message.html',
})
export class Message {
    @Input() message: MessageModel = {
        id: -1,
        author: {
            username: 'undefined author',
            avatarPath: '',
        },
        content: 'undefined content',
    };
}
