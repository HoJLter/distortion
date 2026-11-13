import { Service } from '@angular/core';

@Service()
export class UserService {
    private username: string = 'HoJLter';
    private uid: string = '6cefb5';
    private avatarPath: string = "/"

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
