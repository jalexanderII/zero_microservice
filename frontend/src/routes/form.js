import React, { useState } from 'react';
import { CreateRealtorRequest, Realtor } from '../proto/listings/realtor_pb'
import { Styles, getListingsClient } from "../clients";

const srv = getListingsClient()

export default function CreateRealtor() {
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [phone_number, setPhoneNumber] = useState('')
    const [company, setCompany] = useState('')
    const req = new CreateRealtorRequest()
    const R = new Realtor();
    const handleSubmit = e => {
        e.preventDefault()
        R.setName(name)
        R.setEmail(email)
        R.setPhoneNumber(phone_number)
        R.setCompany(company)
        req.setRealtor(R)
        srv.createRealtor(req, {}).then((resp) => {
            console.log(resp.toObject());
        })
    }
    return (
        <Styles>
        <form onSubmit={handleSubmit}>
            <label htmlFor="name">Enter realtor name</label> <br />
            <input
                name="name"
                value={name}
                onChange={e => setName(e.target.value)}
            />
            <br />
            <label htmlFor="email">Enter realtor Email</label>
            <textarea
                type="email"
                name="email"
                value={email}
                onChange={e => setEmail(e.target.value)}
            />
            <br />
            <label htmlFor="phone_number">Enter realtor number</label> <br />
            <input
                name="phone_number"
                value={phone_number}
                onChange={e => setPhoneNumber(e.target.value)}
            />
            <br />
            <label htmlFor="company">Enter realtor company</label>
            <textarea
                name="company"
                value={company}
                onChange={e => setCompany(e.target.value)}
            />
            <br />
            <button type="submit">Submit!</button>
        </form>
        </Styles>
    )
}
