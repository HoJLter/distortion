import { Component, Input } from '@angular/core';
import { MemberGroup } from './member-group/member-group';
import { UserGroupModel } from '../../shared/data/models/user.models';


@Component({
    imports: [MemberGroup],
    selector: 'server-members',
    styleUrl: './members-sidebar.css',
    templateUrl: './members-sidebar.html',
})
export class MembersSidebar {
    @Input() groups: UserGroupModel[] = [
        {
            name: 'online',
            users: [
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
            users: [
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
