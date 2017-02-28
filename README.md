# Go CONNECT demo client

This is a demo client for Telenor CONNECT ID go library at
 https://github.com/telenordigital/goconnect

## Running the demo
Lunch the demo the usual way:

    go build && ./goconnect-sample

...then point your browser to `http://localhost:8080/` to test it.
Without logging in you can go to `http://localhost:8080/api/userinfo` directly.
You should see an error message. Go back to the root of the service and 
log in by clicking on the button. If you haven't registered you will be sent
through a sign-up. Once th sign-in is complete you should be able to access
the `http://localhost:8080/api/userinfo` from your browser directly. 

## About the development credentials
The development credentials included in this sample works only on a server 
running on your own computer and the browser will be redirected to
http://localhost:8080/connect/oauth2callback when an authentication round-trip is 
completed. Naturally you don't want this for your production services.

The credentials included are of the "public" type which means you can safely 
use these in  clients that do not have a backend to keep the secrets such as 
pure JavaScript services and IOS/Android apps without a dedicated backend 
service for authentication.

