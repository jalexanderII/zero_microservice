import { AuthServiceClient } from './proto/users/auth_grpc_web_pb'
import { ListingsPromiseClient } from './proto/listings/listings_grpc_web_pb'
import styled from "styled-components";

export const Styles = styled.div`
  background: lavender;
  padding: 20px;

  h1 {
    border-bottom: 1px solid white;
    color: #3d3d3d;
    font-family: sans-serif;
    font-size: 20px;
    font-weight: 600;
    line-height: 24px;
    padding: 10px;
    text-align: center;
  }

  form {
    background: white;
    border: 1px solid #dedede;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    margin: 0 auto;
    max-width: 500px;
    padding: 30px 50px;
  }

  input {
    border: 1px solid #d9d9d9;
    border-radius: 4px;
    box-sizing: border-box;
    padding: 10px;
    width: 100%;
  }

  label {
    color: #3d3d3d;
    display: block;
    font-family: sans-serif;
    font-size: 14px;
    font-weight: 500;
    margin-bottom: 5px;
  }

  .error {
    color: red;
    font-family: sans-serif;
    font-size: 12px;
    height: 30px;
  }

  .submitButton {
    background-color: #6976d9;
    color: white;
    font-family: sans-serif;
    font-size: 14px;
    margin: 20px 0px;
  }
`;

export function getAuthServiceClient() {
  const authServiceClient = new AuthServiceClient("http://localhost:8080")
  return authServiceClient;
}

export function getListingsClient() {
  /**
   * @constructor
   * @implements {UnaryInterceptor}
   */
  const SimpleUnaryInterceptor = function() {};

  /** @override */
  SimpleUnaryInterceptor.prototype.intercept = function(request, invoker) {
    console.log('calling SimpleUnaryInterceptor');
    const metadata = request.getMetadata()
    metadata.Authorization = localStorage.getItem("token")
    console.log(request.getMetadata());
    return invoker(request)
  };
  let opts = { 'unaryInterceptors': [new SimpleUnaryInterceptor()] };
  const listingsClient = new ListingsPromiseClient("http://localhost:8080", null, opts)
  return listingsClient;
}
