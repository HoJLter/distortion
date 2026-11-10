import { Service } from '@angular/core';

@Service()
export class UserService {
    private username: string = 'HoJLter';
    private uid: string = '6cefb5';

    public getUsername() {
        return this.username;
    }

    public getUID() {
        return this.uid;
    }
}
