import React, { Component } from 'react';
import logo from '../logo.svg';
import '../App.css';
import { GetApartmentRequest, ListApartmentRequest } from '../proto/listings/apartment_pb'
import { FeedbackForm } from '../form'
import { getListingsClient } from "../clients";

const srv = getListingsClient()

class Listings extends Component  {
  state = {
    resp: null,
    listResult: [],
  }

  GA = () => {
    const req = new GetApartmentRequest()
    req.setId(1)
    srv.getApartment(req, {}, (err, resp) => {
      if (err) {
        console.log(err.code);
        console.log(err.message);
      } else {
        console.log(resp.toObject());
      }
      this.setState({ rent: resp.getRent() })
      this.setState({ fullAddress: resp.getFullAddress() })
    })
  }

  LA = () => {
    const listReq = new ListApartmentRequest()
    srv.listApartments(listReq, {}, (err, result) => {
      if (err) {
        console.log(err.code);
        console.log(err.message);
      } else {
        console.log(result.toObject());
        this.setState({ listResult: result.getApartmentsList() })
      }
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
            <FeedbackForm />
            <button onClick={this.GA}>Get Apartment</button>
              <div>Apartmnet:{this.state.fullAddress} with rent of {this.state.rent}</div>
            <button onClick={this.LA}>Apartment List</button>
            {this.state.listResult.map(i => (<li key="{i.getId()}">{i.getFullAddress()} costs: {i.getRent()}</li>))}
        </header>
      </div>
    );
  }
}

export default Listings;