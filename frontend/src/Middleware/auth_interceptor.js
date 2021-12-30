class AuthInterceptor {
    constructor(token) {
        this.token = token
    }

    intercept(request, invoker) {
        console.log("client side interceptor")
        const metadata = request.getMetadata()
        metadata.Authorization = this.token
        return invoker(request)
    }
}

export {AuthInterceptor};