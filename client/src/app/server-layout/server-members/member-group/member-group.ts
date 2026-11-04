import { Component, Input } from '@angular/core';
import { Member } from '../../../shared/member/member';
import { MemberModel } from '../../../shared/models';

@Component({
    imports: [Member],
    selector: 'member-group',
    styleUrl: './member-group.css',
    templateUrl: './member-group.html',
})
export class MemberGroup {
    @Input() name: string = 'Header';
    @Input() members: MemberModel[] = [];
}
