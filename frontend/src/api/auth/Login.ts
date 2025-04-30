import { LoginURL } from "../Variables";
import { LoginCredentials, TokenResponse } from "../../types/auth";
import axios from "axios";

export const LoginRequest = (request: LoginCredentials): Promise<TokenResponse> => {
  return axios
    .post<TokenResponse>(LoginURL, request)
    .then(res => res.data)
};
