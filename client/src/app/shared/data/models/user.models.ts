export interface UserModel {
    username: string;
}

export interface UserGroupModel {
    name: string;
    users: UserModel[];
}

