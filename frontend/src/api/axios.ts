import axios from "axios";

export const api =  axios.create({
    baseURL: "http://localhost:8080/api/v1",
    timeout: 10_000,
    headers: {
        'Content-Type': "application/json"
    }
})

let getToken = (): null | string => null

export const setAuthTokenGetter = (fn: () => string | null) => {
    getToken = fn;
}

api.interceptors.request.use(config => {
    const token = getToken();
    if (token) config.headers.Authorization = `Bearer ${token}`
    return config
})