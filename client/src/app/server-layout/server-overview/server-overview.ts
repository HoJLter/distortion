import { Component, Input } from '@angular/core';
import { Member } from '../../shared/member/member';
import { MemberGroup } from '../server-members/member-group/member-group';
import { ServerDescription } from './server-description/server-description';
import { ChannelGroup } from './channel-group/channel-group';
import { Profile } from '../../shared/profile/profile';
import { ChannelGroupModel } from '../../shared/models';



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
                    name: 'voice 1',
                    type: 'voice',
                    uid: '',
                    users: [],
                },
                {
                    name: 'voice 2',
                    type: 'voice',
                    uid: '',
                    users: [],
                },
                {
                    name: 'voice 3',
                    type: 'voice',
                    uid: '',
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
                    uid: '',
                    users: [],
                },
                {
                    name: 'text 2',
                    type: 'text',
                    uid: '',
                    users: [],
                },
                {
                    name: 'text 3',
                    type: 'text',
                    uid: '',
                    users: [],
                },
            ],
        },
    ];
}
