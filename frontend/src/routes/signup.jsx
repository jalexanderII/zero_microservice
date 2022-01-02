import React, { useState } from 'react';
import { getAuthServiceClient } from "../clients";
import { Styles } from "../styledComponent";
import { SignupRequest } from '../proto/users/auth_pb'

let authServiceClient = getAuthServiceClient();

let token = ""

export default function SignUp() {
    const [username, setUsername] = useState('')
    const [email, setEmail] = useState('')
    const [role, setRole] = useState('')
    const [password, setPassword] = useState('')

    let req = new SignupRequest()
    const handleSubmit = e => {
        e.preventDefault()
        req.setUsername(username)
        req.setPassword(password)
        req.setEmail(email)
        req.setRole(role)
        authServiceClient.signUp(req, {}, (err, resp) => {
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
            <label htmlFor="email">Email</label> <br />
            <input
                name="email"
                value={email}
                onChange={e => setEmail(e.target.value)}
            />
            <br />
            <label htmlFor="role">Role</label>
            <textarea
                name="role"
                value={role}
                onChange={e => setRole(e.target.value)}
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
