import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { GetApartmentRequest, ListApartmentRequest } from './proto/apartment_pb'
import { ListingsClient } from './proto/listings_grpc_web_pb'
import { ErrorBoundary } from './ErrorBoundary'
import { FeedbackForm } from './form'

const srv = new ListingsClient("http://localhost:8080")

class App extends Component {

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
      <ErrorBoundary>
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
      </ErrorBoundary >
    );
  }
}

export default App;