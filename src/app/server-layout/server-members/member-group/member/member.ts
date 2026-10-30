import { Component, Input } from '@angular/core';

@Component({
  imports: [],
  selector: 'member',
  styleUrl: './member.css',
  templateUrl: './member.html',
})
export class Member {
    @Input() username: string = "username"
}
