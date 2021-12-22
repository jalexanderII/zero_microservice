import * as grpcWeb from 'grpc-web';
import React from 'react';
import logo from './logo.svg';
import './App.css';
// Import the client and the message definition
import { ListingsClient } from './proto/ListingsServiceClientPb';
import { GetApartmentRequest, Apartment } from './proto/apartment_pb';

function App() {
  var echoService = new ListingsClient('http://localhost:8080');

  var request = new GetApartmentRequest();
  request.setId(1);
  var metadata = {};

  const call = echoService.getApartment(request, metadata,
    (err: grpcWeb.RpcError, response: Apartment) => {
      if (err) {
        console.log(err.code);
        console.log(err.message);
      } else {
        console.log(response.toObject());
      }
    });
  call.on('status', (status: grpcWeb.Status) => {
    if (status.metadata) {
      console.log('Received metadata');
      console.log(status.metadata);
    }
  });
  
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
