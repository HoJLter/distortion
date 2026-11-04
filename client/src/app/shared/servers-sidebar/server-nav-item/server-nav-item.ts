import { Component, Input } from '@angular/core';

@Component({
    imports: [],
    selector: 'server-nav-item',
    styleUrl: './server-nav-item.css',
    templateUrl: './server-nav-item.html',
})
export class ServerNavItem {
    @Input() iconPath: string = "";
}

