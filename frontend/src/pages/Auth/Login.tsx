import { Box, Button, TextField, Typography } from "@mui/material";
import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";
import { LoginRequest } from "../../api/auth/Login";
import { ServiceErrorResponse } from "../../types/errors";
import { LoginCredentials } from "../../types/auth";
import { FormEvent } from "react";
import { useSnackbar } from "notistack";

export const LoginPage: React.FC = () => {
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const [credentials, setCredentials] = useState<LoginCredentials>({
    username: "",
    password: "",
  });

  const auth = useContext(AuthContext);

  if (!auth) throw new Error("AuthContext not provided");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  useEffect(() => {
    if (auth.token) navigate("/dashboard");
  });

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const data = await LoginRequest(credentials);
      auth.login(data.token);
    } catch (err: any) {
      let message = "";
      if (err.response && err.response.data) {
        const svcErr = err.response.data as ServiceErrorResponse;
        message = svcErr.message;
      } else {
        message = "Неверные данные или сетевая ошибка";
      }
      enqueueSnackbar(message || "Неизвестная ошибка", { variant: "error" });
    }
  };

  return (
    <Box
      sx={{
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Box
        component="form"
        onSubmit={handleSubmit}
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: "10px",
          width: "25vw",
          border: "1px solid white",
          padding: 2,
        }}
      >
        <Typography variant="h5">Вход</Typography>
        <TextField
          label="Имя пользователя"
          name="username"
          value={credentials.username}
          onChange={handleChange}
          variant="standard"
          required
        />
        <TextField
          label="Пароль"
          name="password"
          type="password"
          value={credentials.password}
          onChange={handleChange}
          variant="standard"
          required
        />
        <Button type="submit" variant="contained">
          Вход
        </Button>
        <Button
          type="button"
          variant="outlined"
          color="primary"
          onClick={() => navigate("/signup")}
        >
          Регистрация
        </Button>
      </Box>
    </Box>
  );
};
