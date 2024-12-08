export type TodoData = {
    id: string;
    title: string;
    isCompleted: boolean;
}

export type CommonResponse = {
    isSuccess: boolean;
    result: any;
}

export type Profile = {
    aud: string;
    exp: string;
    iat: string;
    iss: string;
    name: string;
    nickname: string;
    picture: string;
    sid: string;
    sub: string;
    updated_at: string;
}