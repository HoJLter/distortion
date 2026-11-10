import { Component, Input } from '@angular/core';
import { Member } from '../../shared/member/member';
import { MemberGroup } from '../server-members/member-group/member-group';
import { ServerDescription } from './server-description/server-description';
import { ChannelGroup } from './channel-group/channel-group';
import { Profile } from '../../shared/profile/profile';
import { ChannelGroupModel } from '../../shared/data/models/channel.models';



@Component({
    imports: [ServerDescription, ChannelGroup, Profile],
    selector: 'server-overview',
    styleUrl: './server-overview.css',
    templateUrl: './server-overview.html',
})
export class ServerOverview {
    @Input() serverName: string = 'Server name';

    username: string = 'WASSUP';

    channelGroups: ChannelGroupModel[] = [
        {
            name: 'voice channels',
            type: 'voice',
            channels: [
                {
                    uid: 'voice 1',
                    type: 'voice',
                    users: [],
                },
                {
                    uid: 'voice 2',
                    type: 'voice',
                    users: [],
                },
                {
                    uid: 'voice 3',
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
                    uid: 'text 1',
                    type: 'text',
                    users: [],
                },
                {
                    uid: 'text 2',
                    type: 'text',
                    users: [],
                },
                {
                    uid: 'text 3',
                    type: 'text',
                    users: [],
                },
            ],
        },
    ];
}
