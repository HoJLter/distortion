import { Component, Input } from '@angular/core';
import { Member } from '../../shared/member/member';
import { MemberGroup } from '../server-members/member-group/member-group';
import { ServerDescription } from './server-description/server-description';
import { Profile } from '../../shared/profile/profile';
import { ChannelGroupModel } from '../../shared/data/models/channel.models';
import { ChannelsBlock } from './channels-block/channels-block';



@Component({
    imports: [ServerDescription, Profile, ChannelsBlock],
    selector: 'server-overview',
    styleUrl: './server-overview.css',
    templateUrl: './server-overview.html',
})
export class ServerOverview {
    @Input() serverName: string = 'Server name';

    username: string = 'WASSUP';
}
