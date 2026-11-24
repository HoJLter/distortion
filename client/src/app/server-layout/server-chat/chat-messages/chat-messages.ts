import {
    afterNextRender,
    Component,
    effect,
    ElementRef,
    inject,
    Injector,
    untracked,
    viewChild,
} from '@angular/core';
import { ChatManager } from '../../../shared/services/presentation/chat-manager';
import { Message } from '../../../shared/message/message';

@Component({
    imports: [Message],
    selector: 'chat-messages',
    styleUrl: './chat-messages.css',
    templateUrl: './chat-messages.html',
})
export class ChatMessages {
    chatManager = inject(ChatManager);
    chatMessages = this.chatManager.messages;
    injector = inject(Injector);

    private host = inject<ElementRef<HTMLElement>>(ElementRef)

    constructor() {
        effect(() => {
            this.chatMessages();
            const element = this.host.nativeElement;

            let distanceToBottom = element.scrollHeight - element.clientHeight - element.scrollTop;

            if (distanceToBottom < 100) {
                afterNextRender(() => {
                    element.scrollTop = element.scrollHeight;

                }, {injector: this.injector});
            }
        })
    }
}
