import React from "react";
import { NavLink, Navigate, Route, Routes } from "react-router-dom";
import HomePage from "./pages/HomePage";
import ProductDetailsPage from "./pages/ProductDetailsPage";
import "./styles/app.css";

const App = () => {
  return (
    <>
      <nav className="app-navbar">
        <div className="app-navbar-inner">
          <NavLink to="/" className="app-brand-link">
            Banana Store
          </NavLink>

          <div className="app-nav-links">
            <NavLink
              to="/"
              className={({ isActive }) =>
                isActive ? "app-nav-link app-nav-link-active" : "app-nav-link"
              }
              end
            >
              Home
            </NavLink>
            <NavLink
              to="/products/1"
              className={({ isActive }) =>
                isActive ? "app-nav-link app-nav-link-active" : "app-nav-link"
              }
            >
              Product Page
            </NavLink>
          </div>
        </div>
      </nav>

      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/products/:id" element={<ProductDetailsPage />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </>
  );
};

export default App;