import { Component, Input } from '@angular/core';
import { UserModel } from '../models';

@Component({
    imports: [],
    selector: 'message',
    styleUrl: './message.css',
    templateUrl: './message.html',
})
export class Message {
    @Input() user: UserModel = {
        name: "Undefined user"
    };
    @Input() message: string = 'Undefined message';
}
