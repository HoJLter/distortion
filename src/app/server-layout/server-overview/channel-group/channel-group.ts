import { Component, Input } from '@angular/core';
import { ChatChannel } from './chat-channel/chat-channel';
import { ChannelModel } from '../../../shared/models';


@Component({
    imports: [ChatChannel],
    selector: 'channel-group',
    styleUrl: './channel-group.css',
    templateUrl: './channel-group.html',
})
export class ChannelGroup {
    @Input() name: string = "placeholder";
    @Input() type: string = "text";
    @Input() channels: ChannelModel[] = [];
}
