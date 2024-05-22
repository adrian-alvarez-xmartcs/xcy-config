import { createContext, useContext, ReactNode, useState } from "react";
import { User } from "../entities/Auth";

interface UserContextType extends User {
    Login: (user:string, password:string) => void
}

const UserContext = createContext<UserContextType | undefined>(undefined)

export const useUser = () => {
    const context = useContext(UserContext);
    if (!context) {
        throw new Error("useUser must be used within a UserProvider")
    }
    return context
}

const INITIAL_STATE:User = {
    name: "",
}

const UserProviderInternal = ({children}: {children: ReactNode}) => {
    const [user, setUser] = useState<User>(INITIAL_STATE)

    const Login = (user:string, password:string) => {
        console.log("asdas")
        setUser(prevState => ({
            ...prevState,
            name: user + password
        }))
    }

    return (
        <UserContext.Provider value={{...user, Login}}>
            {children}
        </UserContext.Provider>
    ) 
}

export const UserProvider = ({children}: {children: ReactNode}) => {
    return (
        <UserProviderInternal>
            {children}
        </UserProviderInternal>
    )
}