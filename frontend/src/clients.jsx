import { AuthServiceClient } from './proto/users/auth_grpc_web_pb'
import { ListingsPromiseClient } from './proto/listings/listings_grpc_web_pb'

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
