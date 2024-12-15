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
    name?: string;
    nickname?: string;
    family_name?: string;
    middle_name?: string;
    picture?: string;
    email?: string;
    email_verified?: boolean;
    gender?: string;
    birthdate?: string;
    phone_number?: string;
    phone_number_verified?: boolean;
    updated_at?: string;
}