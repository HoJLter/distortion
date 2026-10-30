import { Component, Input } from '@angular/core';
import { MemberModel } from '../../server-members/member-group/member-group';
import { ChatChannel } from './chat-channel/chat-channel';

export interface ChannelModel {
    name: string;
    type: string;
    users: MemberModel[];
}

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
