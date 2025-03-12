import { AppBar, Box } from '@mui/material';
import React from 'react';

export const HomePage = () =>  {
    return (
        <Box>
            <Box sx={{flexGrow: 1}}>
                <AppBar position='static'>
                </AppBar>
            </Box>
            <h1>Home Page</h1>
        </Box>
    );
}