import { Component, Input } from '@angular/core';
import { Member } from '../server-members/member-group/member/member';
import { MemberGroup, MemberModel } from '../server-members/member-group/member-group';
import { ServerDescription } from './server-description/server-description';
import { ChannelGroup, ChannelModel } from './channel-group/channel-group';

export interface ChannelGroupModel {
    name: string;
    type: string;
    channels: ChannelModel[];
}

@Component({
    imports: [MemberGroup, ServerDescription, ChannelGroup],
    selector: 'server-overview',
    styleUrl: './server-overview.css',
    templateUrl: './server-overview.html',
})
export class ServerOverview {
    @Input() serverName: string = "Server name";

    channelGroups: ChannelGroupModel[] = [
        {
            name: 'voice channels',
            type: 'voice',
            channels: [
                {
                    name: 'voice 1',
                    type: 'voice',
                    users: [],
                },
                {
                    name: 'voice 2',
                    type: 'voice',
                    users: [],
                },
                {
                    name: 'voice 3',
                    type: 'voice',
                    users: [],

                },
            ],
        },
        {
            name: 'text channels',
            type: 'text',
            channels: [
                {
                    name: 'text 1',
                    type: 'text',
                    users: [],
                },
                {
                    name: 'text 2',
                    type: 'text',
                    users: [],
                },
                {
                    name: 'text 3',
                    type: 'text',
                    users: [],
                },
            ],
        }
    ];
}
