import { Profile } from "../models/types";
import { AUTH_URL } from "../utils/constants";
import { getMethodGetHeader } from "../utils/https";

export async function login(): Promise<{ url: string }> {
    const headers = getMethodGetHeader()
    const response = await fetch(AUTH_URL, headers);
    const data = await response.json()
    return data;
}

export async function callback(code: string): Promise<Profile> {
    const url = AUTH_URL + "/callback/" + code;
    const headers = getMethodGetHeader()
    const response = await fetch(url, headers);
    const data = await response.json() as Profile;
    return data;
}