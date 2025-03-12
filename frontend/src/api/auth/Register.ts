import { ServiceErrorResponse } from "../Types";
import { RegisterURL } from "../Variables"
import { LoginRequest, TokenResponse } from "./Types"

export const Request = async (request: LoginRequest) => {

    let result = await (await fetch(RegisterURL)).json(); 
    
    if (typeof result?.code === "number") {
        return result as ServiceErrorResponse
    }

    return (result as TokenResponse).token
}