import React from 'react';
import { Home } from './views/Home';
import { NavBar } from './components/NavBar';
import { Route, Routes, Navigate } from 'react-router-dom';

export const RouteLinks = () => {
    return (
    <div>
        <NavBar/>
        <Routes>
            <Route path="/Home" component={Home} />
            <Route path="/">
                <Navigate to="/Home" />
            </Route>
        </Routes>
    </div>
  );
};