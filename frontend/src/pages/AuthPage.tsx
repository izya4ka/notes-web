import { Box, Button, Container, TextField } from "@mui/material";

export const AuthPage = () => {
    return (
        <Box sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            flexDirection: "column",
            height: "70vh"
        }}>
            <Container sx={{
                width: "50%",
                height: "50%",
                border: "1px solid",
                display: "flex",
                flexDirection: "column"
            }}>
                <h1>Sign in</h1>
                <TextField
                    required
                    id="standard"
                    label="Login"
                    variant="standard"
                />
                <TextField
                    required
                    id="standard"
                    label="Password"
                    variant="standard"
                    type="password"
                    sx={{marginTop: "5px"}}
                />
                <Box marginTop={"20px"}>
                    <Button variant="outlined" sx={{maxWidth: "20%"}}>Login</Button>
                    <Button variant="text" sx={{marginLeft: "15px"}}>Sign up</Button>
                </Box>
            </Container>
        </Box>
    );
}