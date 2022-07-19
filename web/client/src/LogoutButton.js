import React from 'react'

const LogoutButton = (callback) => {
    return (
        <input type="button" onClick={callback} value="Logout"></input>
    )
}

export default LogoutButton;