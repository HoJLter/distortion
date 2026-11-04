import { Component, Input } from '@angular/core';
import { MemberGroup } from './member-group/member-group';
import { MemberGroupModel } from '../../shared/models';


@Component({
    imports: [MemberGroup],
    selector: 'server-members',
    styleUrl: './members-sidebar.css',
    templateUrl: './members-sidebar.html',
})
export class MembersSidebar {
    @Input() groups: MemberGroupModel[] = [
        {
            name: 'online',
            members: [
                {
                    username: 'hojlter',
                },
                {
                    username: 'hojlter',
                },
                {
                    username: 'hojlter',
                },
                {
                    username: 'hojlter',
                },
                {
                    username: 'hojlter',
                },
            ],
        },
        {
            name: 'offline',
            members: [
                {
                    username: 'megalul',
                },
                {
                    username: 'megalul',
                },
                {
                    username: 'megalul',
                },
                {
                    username: 'megalul',
                },
            ],
        },
    ];
}
