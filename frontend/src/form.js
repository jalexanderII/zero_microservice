import React, { useState } from 'react';
import { CreateRealtorRequest, Realtor } from './proto/realtor_pb'
import { ListingsClient } from './proto/listings_grpc_web_pb'

var FeedbackForm = function () {
    const [id, setId] = useState('')
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [phone_number, setPhoneNumber] = useState('')
    const [company, setCompany] = useState('')
    const srv = new ListingsClient("http://localhost:8080")
    const req = new CreateRealtorRequest()
    const R = new Realtor();
    const handleSubmit = e => {
        e.preventDefault()
        R.setId(id)
        R.setName(name)
        R.setEmail(email)
        R.setPhoneNumber(phone_number)
        R.setCompany(company)
        req.setRealtor(R)
        srv.createRealtor(req, {}, (err, resp) => {
            if (err) {
                console.log(err.code);
                console.log(err.message);
            } else {
                console.log(resp.toObject());
            }
        })
    }
    return (
        <form onSubmit={handleSubmit}>
            <label htmlFor="id">Enter realtor id</label>
            <textarea
                name="id"
                value={id}
                onChange={e => setId(e.target.value)}
            />
            <br />
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
    )
}

export {FeedbackForm};