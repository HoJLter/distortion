import { Component, Input } from '@angular/core';

@Component({
    imports: [],
    selector: 'chat-header',
    styleUrl: './chat-header.css',
    templateUrl: './chat-header.html',
})
export class ChatHeader {
    @Input() header: string = "Undefined header";
}
