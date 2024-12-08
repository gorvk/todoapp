export const getMethodGetHeader = (): RequestInit | undefined => {
    return {
        method: "GET",
        credentials: "include"
    }
}

// eslint-disable-next-line
export const getMethodPostHeader = (payload: any): RequestInit | undefined => {
    return {
        method: "POST",
        body: JSON.stringify(payload),
        credentials: "include"
    }
}

// eslint-disable-next-line
export const getMethodPutHeader = (payload: any): RequestInit | undefined => {
    return {
        method: "PUT",
        body: JSON.stringify(payload),
        credentials: "include"
    }
}