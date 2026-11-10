import { Component, Input } from '@angular/core';
import { UserModel } from '../data/models/user.models';

@Component({
    imports: [],
    selector: 'message',
    styleUrl: './message.css',
    templateUrl: './message.html',
})
export class Message {
    @Input() user: UserModel = {
        username: "Undefined user"
    };
    @Input() message: string = 'Undefined message';
}
