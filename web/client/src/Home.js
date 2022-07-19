import React from "react";
import LogoutButton from "./LogoutButton";
import { useNavigate } from "react-router-dom";
import LoginContext from "./contexts/LoginContext";

const Home = () => {
  const {isLogin, setLogin} = React.useContext(LoginContext);
  const navigate = useNavigate()
  const doLogout = () => {
    setLogin(false);
    navigate("/");
  };

  return (
    <div>
      <h1>Home</h1>
      <LogoutButton callback={doLogout}/>
    </div>
  );
};

export default Home;
