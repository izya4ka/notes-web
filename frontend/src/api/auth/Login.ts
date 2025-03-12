import { ServiceErrorResponse } from "../Types";
import { LoginURL } from "../Variables"
import { LoginRequest, TokenResponse } from "./Types"

export const Login = async (request: LoginRequest) => {

    let result = await (await fetch(LoginURL)).json(); 
    
    if (typeof result?.code === "number") {
        return result as ServiceErrorResponse
    }

    return (result as TokenResponse).token
}