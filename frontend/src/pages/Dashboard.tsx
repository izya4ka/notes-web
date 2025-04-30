import { Box } from "@mui/material";
import React from "react";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { Navbar } from "../components/Navbar";

export const Dashboard: React.FC = () => {
  const auth = useContext(AuthContext);
  if (!auth) throw new Error("AuthContext not provided");

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        width: "100vw",
        height: "100vh",
      }}
    >
      <Navbar />
    </Box>
  );
};
