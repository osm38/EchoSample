import React from 'react'
import { useNavigate } from 'react-router';

const LoginButton = ({callback}) => {

    return (
        // <input type="button" onClick={() => setLogin(true)}>Login</input>
        // <button type="submit">Login</button>
        <input type="button" value="Login" onClick={callback} />
    )
}

export default LoginButton;