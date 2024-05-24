import { mockUsers, mockWorkspaces } from './mockData';

export const mockLogin = (username: string, password: string, workspace: string) => {
    console.log(mockUsers)
    const foundUser = mockUsers.find(
        mockUser =>
            mockUser.username === username &&
            mockUser.password === password &&
            mockUser.workspace === workspace
    );
    console.log(foundUser)
    if (foundUser) {
        console.log(foundUser)
        return { ok: true, data: { username: foundUser.username, workspace: foundUser.workspace } };
    } else {
        console.log("bad login")
        return { ok: false, message: "Invalid credentials" };
    }
};

export const mockGetWorkspaces = () => {
    return { ok: true, data: mockWorkspaces };
};
