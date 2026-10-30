import { Component, Input } from '@angular/core';

@Component({
    imports: [],
    selector: 'server-description',
    styleUrl: './server-description.css',
    templateUrl: './server-description.html',
})
export class ServerDescription {
    @Input() serverName: string = "";
}
