import { useUser } from "../contexts/UserContext"

const LoginPage = () => {
    const userCtx = useUser()

    return (
        <>
            <p>{userCtx.name}</p>
            <button onClick={() => userCtx.Login("sdadas", "sdasda2")}>
                LOGIN
            </button>
        </>
    )
}

export default LoginPage