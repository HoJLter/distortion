import { Component, computed, effect, inject, Input } from '@angular/core';
import { MemberGroup } from './member-group/member-group';
import { UserGroupModel } from '../../shared/data/models/user.models';
import {UserService} from '../../shared/services/user-service';


@Component({
    imports: [MemberGroup],
    selector: 'server-members',
    styleUrl: './members-sidebar.css',
    templateUrl: './members-sidebar.html',
})
export class MembersSidebar {
    userService: UserService = inject(UserService);

    groups = computed<UserGroupModel[]>(() => [
        {
            name: "Users",
            users: this.userService.users()
        }
    ])

}
