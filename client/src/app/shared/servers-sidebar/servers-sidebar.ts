import { Component, Input } from '@angular/core';
import { ServerNavItemModel } from '../data/models/server.models';
import { ServerNavItem } from './server-nav-item/server-nav-item';


@Component({
    imports: [ServerNavItem],
    selector: 'servers-sidebar',
    styleUrl: './servers-sidebar.css',
    templateUrl: './servers-sidebar.html',
})
export class ServersSidebar {
    @Input() servers: ServerNavItemModel[] = [];
}
