import { RegisterURL } from "../Variables";
import { LoginCredentials, TokenResponse } from "../../types/auth";
import axios from "axios";

export const RegisterRequest = (
  request: LoginCredentials
): Promise<TokenResponse> => {
  return axios
    .post<TokenResponse>(RegisterURL, request)
    .then((res) => res.data);
};
