
export interface GetUserResponse {
    user_uid: string;
    username: string;
    avatar_path: string;
}

export interface CreateUserRequest {
    username: string;
    avatar_path: string;
}
