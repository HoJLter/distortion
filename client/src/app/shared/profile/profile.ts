import { Component, Input } from '@angular/core';

@Component({
    imports: [],
    selector: 'profile',
    styleUrl: './profile.css',
    templateUrl: './profile.html',
})
export class Profile {
    @Input() username: string = 'username';
}
