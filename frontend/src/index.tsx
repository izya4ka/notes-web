import React from "react";
import ReactDOM from "react-dom/client";
import { App } from "./App";
import { ThemeProvider } from "@emotion/react";
import { CssBaseline } from "@mui/material";
import { themeOptions } from "./style/theme";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <ThemeProvider theme={themeOptions}>
    <CssBaseline />
    <App />
  </ThemeProvider>
);
