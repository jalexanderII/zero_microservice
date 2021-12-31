import React, { Component } from 'react';
import logo from '../logo.svg';
import '../App.css';
import { GetApartmentRequest } from '../proto/listings/apartment_pb'
import { getListingsClient } from "../clients";

const srv = getListingsClient()

class Listings extends Component  {
  constructor(props) {
    super(props);
    this.state = {
      resp: null,
      value: '',
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  GA = () => {
    const req = new GetApartmentRequest()
    req.setId(this.state.value)
    srv.getApartment(req, {}).then((resp) => {
      console.log(resp.toObject());
      this.setState({ rent: resp.getRent() })
      this.setState({ fullAddress: resp.getFullAddress() })
    })
  }

  handleSubmit(event) {
    this.GA()
    event.preventDefault();
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <form onSubmit={this.handleSubmit}>
          <label>
            Apartment ID:
          <input type="text" value={this.state.value} onChange={this.handleChange} />
          </label>
          <input type="submit" value="Submit" />
          </form>
              {this.state.rent > 0 && <h3>Apartmnet:{this.state.fullAddress} with rent of {this.state.rent}</h3> }
        </header>
      </div>
    );
  }
}

export default Listings;