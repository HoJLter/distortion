import { Service, signal, Signal, WritableSignal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UserModel } from '../data/models/user.models';

@Service()
export class UserService {
    private username: string = 'HoJLter';
    private uid: string = '6cefb5';
    private avatarPath: string = "/"

    public users: WritableSignal<UserModel[]> = signal<UserModel[]>([])

    constructor() {
        
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
