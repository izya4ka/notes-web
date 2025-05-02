import { LoginCredentials, TokenResponse } from "../../types/auth";
import { ServiceErrorResponse } from "../../types/errors";
import { api } from "../axios";

export const RegisterRequest = (
  request: LoginCredentials
): Promise<TokenResponse | ServiceErrorResponse> => {
  return api
    .post<TokenResponse | ServiceErrorResponse>("/user/register", request)
    .then((res) => res.data);
};
