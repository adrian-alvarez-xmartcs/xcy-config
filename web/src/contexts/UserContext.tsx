import { createContext, useContext, ReactNode, useState } from "react";
import { ExtendedUser } from "../entities/Auth";
import ApiBackend from "../apis/ApiBackend";

interface UserContextType extends ExtendedUser {
    Login: (username: string, password: string, workspace: string) => Promise<boolean>;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const useUser = () => {
    const context = useContext(UserContext);
    if (!context) {
        throw new Error("useUser must be used within a UserProvider");
    }
    return context;
}

const INITIAL_STATE: ExtendedUser = {
    name: "",
    workspace: ""
};

const UserProviderInternal = ({ children }: { children: ReactNode }) => {
    const [user, setUser] = useState<ExtendedUser>(INITIAL_STATE);

    const Login = async (username: string, password: string, workspace: string): Promise<boolean> => {
        try {
            const response = await ApiBackend.Auth.Login({ username, password, workspace });
            if (response.ok) {
                setUser({ name: username, workspace });
                return true;
            } else {
                return false;
            }
        } catch (error) {
            console.error("Login error:", error);
            return false;
        }
    };

    return (
        <UserContext.Provider value={{ ...user, Login }}>
            {children}
        </UserContext.Provider>
    );
}

export const UserProvider = ({ children }: { children: ReactNode }) => {
    return (
        <UserProviderInternal>
            {children}
        </UserProviderInternal>
    );
}