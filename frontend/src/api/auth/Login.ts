import { LoginCredentials, TokenResponse } from "../../types/auth";
import { ServiceErrorResponse } from "../../types/errors";
import { api } from "../axios";

export const LoginRequest = (
  request: LoginCredentials
): Promise<TokenResponse | ServiceErrorResponse> => {
  return api.post<TokenResponse | ServiceErrorResponse>("/user/login", request).then((res) => res.data);
};
