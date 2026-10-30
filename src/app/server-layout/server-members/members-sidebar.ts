import { Component, Input } from '@angular/core';
import { MemberGroup, MemberModel } from './member-group/member-group';



export interface MemberGroupModel {
    name: string;
    members: MemberModel[];
}

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
