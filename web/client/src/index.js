import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import Login from "./Login";
// import App from './App';
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Dummy from "./Dummy";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <Login></Login>
    <Router>
      <Routes>
        <Route path="/dummy" element={<Dummy value="arg" />}></Route>
      </Routes>
    </Router>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
