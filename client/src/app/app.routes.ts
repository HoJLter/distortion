import { Routes } from '@angular/router';
import { ServerLayout } from './server-layout/server-layout';
import { AuthLayout } from './auth-layout/auth-layout';
import { CreateServerForm } from './auth-layout/create-server-form/create-server-form';
import { JoinServerForm } from './auth-layout/join-server-form/join-server-form';
import { authGuard } from './shared/services/auth/auth.guard';

export const routes: Routes = [
    { path: 'server', component: ServerLayout, canActivate: [authGuard]},
    { path: 'auth', component: AuthLayout, children: [
        { path: '', redirectTo: 'join-server', pathMatch: 'full'},
        { path: 'join-server', component: JoinServerForm },
        { path: 'create-server', component:  CreateServerForm },]},
];
