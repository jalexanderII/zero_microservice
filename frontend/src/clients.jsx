import { AuthServiceClient } from './proto/users/auth_grpc_web_pb'
import { AuthInterceptor } from './Middleware/auth_interceptor'
import { getTokenFromLogin } from './routes/login'
import { getTokenFromSignUp } from './routes/signup'
import { ListingsClient } from './proto/listings/listings_grpc_web_pb'



export function getAuthServiceClient() {
  const authServiceClient = new AuthServiceClient("http://localhost:8080")
  return authServiceClient;
}

export function getListingsClient() {
  let token = getTokenFromLogin()
  if (token === "") {
    token = getTokenFromSignUp()
  }

  const authInterceptor = new AuthInterceptor(token)

  const listingsClient = new ListingsClient("http://localhost:8080", null, { unaryInterceptors: [authInterceptor] })
  return listingsClient;
}
