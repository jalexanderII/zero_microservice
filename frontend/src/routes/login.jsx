import React, { useState } from 'react';
import { Styles, getAuthServiceClient } from "../clients";
import { LoginRequest } from '../proto/users/auth_pb'

let authServiceClient = getAuthServiceClient();

let token = ""

export default function Login() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')

    let req = new LoginRequest()
    const handleSubmit = e => {
        e.preventDefault()
        req.setUsername(username)
        req.setPassword(password)
        authServiceClient.login(req, {}, (err, resp) => {
            if (err) {
                console.log(err.code);
                console.log(err.message);
            } else {
                console.log(resp.toObject());
                token = resp.getToken()
                localStorage.setItem('token', token)
            }
        })
    }
    return (
        <Styles>
        <form onSubmit={handleSubmit}>
            <label htmlFor="username">Username / E-Mail</label>
            <textarea
                name="username"
                value={username}
                onChange={e => setUsername(e.target.value)}
            />
            <br />
            <label htmlFor="password">Password</label> <br />
            <input
                name="password"
                value={password}
                onChange={e => setPassword(e.target.value)}
            />
            <br />
            <button type="submit">Submit!</button>
        </form>
        </Styles>
    )
}
