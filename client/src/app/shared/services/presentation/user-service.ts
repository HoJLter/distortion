import { inject, Service, signal, Signal, WritableSignal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserModel } from '../../data/models/user.models';
import { DbApiService } from '../transport/db-api-service';
import { map } from 'rxjs';
import { mapUser } from '../../data/mappers';

@Service()
export class UserService {
    private username: string = 'HoJLter';
    private uid: string = '6cefb5';
    private avatarPath: string = '/';

    private dbApi = inject(DbApiService);

    public users: WritableSignal<UserModel[]> = signal<UserModel[]>([]);



    constructor() {
        this.dbApi.getUsers().pipe(
            map(usr => usr.map(mapUser)),
        ).subscribe((newUsers) => {
            this.users.update((users) => [...users, ...newUsers])
        });
    }

    public getUsername() {
        return this.username;
    }

    public getUID() {
        return this.uid;
    }

    public getAvatarPath() {
        return this.avatarPath;
    }
}
