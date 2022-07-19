import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Home from "./Home";
import LoginContext from "./contexts/LoginContext";

const Layout = () => {
  const [isLogin, setLogin] = useState(false);
  const value = {
    isLogin,
    setLogin,
  };
  console.log("call Layout!!");

  return (
    <LoginContext.Provider value={value}>
      <Router>
        <Routes>
          {!isLogin && <Route exact path="/" element={<Login />}></Route>}
          {isLogin && <Route path="/home" element={<Home />}></Route>}
        </Routes>
      </Router>
    </LoginContext.Provider>
  );
};

export default Layout;
