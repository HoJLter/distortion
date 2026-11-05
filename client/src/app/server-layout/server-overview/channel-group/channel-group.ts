import { Component, Input, output } from '@angular/core';
import { ChatChannel } from './chat-channel/chat-channel';
import { ChannelGroupModel, ChannelModel } from '../../../shared/models';


@Component({
    imports: [ChatChannel],
    selector: 'channel-group',
    styleUrl: './channel-group.css',
    templateUrl: './channel-group.html',
})
export class ChannelGroup {
    @Input({required: true}) channelGroup!: ChannelGroupModel;
}
