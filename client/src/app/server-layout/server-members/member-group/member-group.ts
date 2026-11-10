import { Component, Input } from '@angular/core';
import { Member } from '../../../shared/member/member';
import { UserGroupModel } from '../../../shared/data/models/user.models';

@Component({
    imports: [Member],
    selector: 'member-group',
    styleUrl: './member-group.css',
    templateUrl: './member-group.html',
})
export class MemberGroup {
    @Input() name: string = 'Header';
    @Input() usersGroup: UserGroupModel = {
        name: "undefined",
        users: []
    };
}
