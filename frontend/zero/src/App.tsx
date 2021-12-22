import * as grpcWeb from 'grpc-web';
import * as React from 'react';
import './App.css';
import { ListingsClient } from './proto/ListingsServiceClientPb';
import { GetApartmentRequest, Apartment } from './proto/apartment_pb';
import { BrowserRouter as Router } from 'react-router-dom';
import { RouteLinks } from './routes'; 


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
    <Router>
      <RouteLinks />
    </Router>
  );
}

export default App;
