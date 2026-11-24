import { Service } from '@angular/core';

@Service()
export class AuthService {

    public isAuthenticated(): boolean {
        return true
    }
}
