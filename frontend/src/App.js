import React from 'react';
import './App.css';
import { ErrorBoundary } from './ErrorBoundary'
import { Outlet, Link } from "react-router-dom";

export default function App() {
  return (
    <ErrorBoundary>
    <div>
      <h1>Zero</h1>
      <nav
        style={{
          borderBottom: "solid 1px",
          paddingBottom: "1rem"
        }}
      >
        <Link to="/listings">Listings</Link> | {" "}
          <Link to="/apartments">Get Apartments</Link> | {" "}
          <Link to="/form">Create Realtor</Link> | {" "}
          <Link to="/login">Login</Link> | {" "}
          <Link to="/signup">Sign Up</Link>
      </nav>
      <Outlet />
    </div>
    </ErrorBoundary >
  );
}