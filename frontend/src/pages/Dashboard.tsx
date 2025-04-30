import { AppBar, Box } from "@mui/material";
import React from "react";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";

export const Dashboard: React.FC = () => {
  const auth = useContext(AuthContext);
  if (!auth) throw new Error('AuthContext not provided');

  return (
    <div>
      <h1>Панель управления</h1>
      <button onClick={auth.logout}>Выйти</button>
    </div>
  );
};