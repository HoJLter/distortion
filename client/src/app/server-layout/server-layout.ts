import { Component } from '@angular/core';
import { ServersSidebar } from '../shared/servers-sidebar/servers-sidebar';
import { ServerChat } from './server-chat/server-chat';
import { ServerOverview } from './server-overview/server-overview';
import { MembersSidebar } from './server-members/members-sidebar';
import { ServerNavModel } from '../shared/models';

@Component({
    imports: [ServersSidebar, ServerChat, ServerOverview, MembersSidebar],
    selector: 'app-server-layout',
    styleUrl: './server-layout.css',
    templateUrl: './server-layout.html',
})
export class ServerLayout {
    serverName: string = 'kuharka';

    servers: ServerNavModel[] = [
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
        {
            iconPath: 'icon',
        },
    ];
}
