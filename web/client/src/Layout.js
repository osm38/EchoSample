// import React from "react";
// import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
// import Login from "./Login";
// import Home from "./Home";
// import LoginContext from "./contexts/LoginContext";

// const Layout = () => {
//   const [isLogin, setLogin] = React.useState(false);
//   const value = {
//     isLogin,
//     setLogin,
//   };

//   return (
//     <LoginContext.Provider value={value}>
//       <Router>
//         <Routes>
//           {isLogin && <Route path="/" element={<Login />}></Route>}
//           {!isLogin && <Route path="/home" element={<Home />}></Route>}
//           <Route path="/" element={<Login />}></Route>
//           <Route path="/home" element={<Home />}></Route>
//         </Routes>
//       </Router>
//     </LoginContext.Provider>
//   );
// };

// export default Layout;

import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Dummy from "./Dummy";
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
          {isLogin && <Route path="/dummy" element={<Dummy value="arg" />}></Route>}
        </Routes>
      </Router>
    </LoginContext.Provider>
  );
};

export default Layout;
