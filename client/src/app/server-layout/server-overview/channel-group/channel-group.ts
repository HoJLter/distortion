import { Component, Input, output } from '@angular/core';
import { ChatChannel } from './chat-channel/chat-channel';
import {ChannelGroupModel} from '../../../shared/data/models/channel.models';


@Component({
    imports: [ChatChannel],
    selector: 'channel-group',
    styleUrl: './channel-group.css',
    templateUrl: './channel-group.html',
})
export class ChannelGroup {
    @Input({required: true}) channelGroup!: ChannelGroupModel;
}
