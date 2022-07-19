import React from 'react'
import LoginContext from './contexts/LoginContext'

const LogoutButton = () => {
    const {isLogin, setLogin} = React.useContext(LoginContext);

    const doLogout = () => {
        console.log("Logout!");
        setLogin(false)
    }

    return (
        <input type="button" onClick={doLogout} value="Logout"></input>
    )
}

export default LogoutButton;