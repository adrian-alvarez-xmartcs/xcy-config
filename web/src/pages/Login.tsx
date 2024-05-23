import React, { useState, useEffect } from 'react';
// import { useUser } from "../contexts/UserContext";
import { TextField, Button, Container, Box, Typography, Autocomplete } from '@mui/material';
import ApiBackend from "../apis/ApiBackend";
interface Workspace {
    Id: number;
    Name: string;
}

const LoginPage = () => {
    // const userCtx = useUser();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [loginAttempted, setLoginAttempted] = useState(false);
    const [workspaces, setWorkspaces] = useState<Workspace[]>([]);
    const [selectedWorkspace, setSelectedWorkspace] = useState<Workspace | null>(null);

    useEffect(() => {
        const fetchWorkspaces = async () => {
            const response = await ApiBackend.Workspace.Get()
            console.log(response.data)
            setWorkspaces(response.data);
        };

        fetchWorkspaces();
    }, []);

    const handleLogin = (event: React.FormEvent) => {
        event.preventDefault();
        setLoginAttempted(true);
    };

    return (
        <Container maxWidth="sm">
            <Box
                component="form"
                onSubmit={handleLogin}
                sx={{
                    px: 4,
                    py: 4,
                    bgColor: 'background.paper',
                    borderRadius: 2,
                    boxShadow: 3,
                    mt: 8,
                }}
            >
                <Typography variant="h4" component="h1" gutterBottom>
                    Xcylla
                </Typography>
                <Box py={2}>
                <Autocomplete
                    options={workspaces}
                    getOptionLabel={(option) => option.Name}
                    value={selectedWorkspace}
                    onChange={(event, newValue) => {
                        setSelectedWorkspace(newValue);
                    }}
                    renderInput={(params) => (
                        <TextField
                            {...params}
                            label="Select Workspace"
                            variant="outlined"
                            fullWidth
                            required
                        />
                    )}
                />
                </Box>
                <Box py={2}>
                    <TextField
                        label="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                        fullWidth
                        variant="outlined"
                        error={loginAttempted && username.length === 0}
                        helperText={loginAttempted && username.length === 0 ? 'Please type something' : ''}
                    />
                </Box>
                <Box pb={2}>
                    <TextField
                        label="Password"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        fullWidth
                        variant="outlined"
                    />
                </Box>
                <Box py={2}>
                    <Button
                        fullWidth
                        size="large"
                        variant="contained"
                        color="primary"
                        type="submit"
                        sx={{ minWidth: 300 }}
                    >
                        Login
                    </Button>
                </Box>
            </Box>
        </Container>
    );
};

export default LoginPage;
