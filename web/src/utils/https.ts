export const getMethodGetHeader = (): RequestInit => {
    return {
        method: "GET",
        headers: {
            'Authorization': `${localStorage.getItem('id_token')}`
        }
    }
}

export const getMethodPostHeader = (payload: any): RequestInit => {
    return {
        method: "POST",
        body: JSON.stringify(payload),
        headers: {
            'Authorization': `${localStorage.getItem('id_token')}`
        }
    }
}

export const getMethodPutHeader = (payload: any): RequestInit => {
    return {
        method: "PUT",
        body: JSON.stringify(payload),
        headers: {
            'Authorization': `${localStorage.getItem('id_token')}`
        }
    }
}

export const getMethodDeleteHeader = (): RequestInit => {
    return {
        method: "DELETE",
        headers: {
            'Authorization': `${localStorage.getItem('id_token')}`
        }
    }
}