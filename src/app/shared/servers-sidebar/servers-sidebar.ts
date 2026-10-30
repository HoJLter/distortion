import { Component, Input } from '@angular/core';
import { ServerNavItem } from './server-nav-item/server-nav-item';


export interface ServerNavModel {
    iconPath: string;
}

@Component({
    imports: [ServerNavItem],
    selector: 'servers-sidebar',
    styleUrl: './servers-sidebar.css',
    templateUrl: './servers-sidebar.html',
})
export class ServersSidebar {
    @Input() servers: ServerNavModel[] = [];
}
