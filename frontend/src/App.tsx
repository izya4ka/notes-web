import React, { useMemo } from 'react';
import {
  BrowserRouter,
  Route,
  Routes
} from "react-router-dom";
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { AuthPage } from './pages/AuthPage';
import { HomePage } from './pages/HomePage';
import { createTheme, CssBaseline, ThemeProvider } from '@mui/material';

export const App = () => {

  const theme = useMemo(() => createTheme({
     palette: { mode: "dark" },
    }), []);
    
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<AuthPage />} />
          <Route path="home" element={<HomePage />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}
