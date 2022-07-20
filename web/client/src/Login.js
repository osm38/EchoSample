// import { Link } from "react-router-dom"
import LoginButton from './LoginButton'
import React from 'react';
import { useNavigate } from 'react-router-dom';
import LoginContext from './contexts/LoginContext'

const Login = () => {
  const {isLogin, setLogin} = React.useContext(LoginContext);
  const navigate = useNavigate();

  const doLogin = () => {
      console.log(isLogin);
      setLogin(true);
      navigate("/home");
  }

  return (
    <form className="w-[400px] px-2.5 py-10 absolute text-center top-2/4 left-2/4 -translate-x-2/4 -translate-y-2/4 bg-black">
      <div>
        <h1 className="text-white font-medium">LOGIN</h1>
      </div>
      <div className="py-2.5 bg-gray-200 my-2.5">
        <input
          className="w-50 p-1 my-2"
          type="id"
          name="name"
          placeholder="Input Your ID"
        />
        <br />
        <input
          className="w-50 p-1 my-2"
          type="password"
          name="password"
          placeholder="Input Your Password"
        />
      </div>
      {/* <button type="submit">Login</button> */}
      <LoginButton callback={doLogin}/>
    </form>
  );
}

export default Login;
