# go-code-together
A collaborative web based plain text editor with a server written in Go.

My first project in Golang!

Translated from my Node version [here](https://github.com/dyang108/webcoder). Using [socket.io](socket.io), [go-socket.io](https://github.com/googollee/go-socket.io), [mgo](https://github.com/go-mgo/mgo), [Ace text editor](https://ace.c9.io/) and Bootstrap.

## Build

### Install [Go](https://golang.org/doc/install)
Now install Go packages: 

    go get -d ./...

### Clone repository somewhere in your $GOPATH

    git clone https://github.com/dyang108/go-code-together.git
    cd go-code-together

### Install [npm](https://nodejs.org/en/)
Now install node modules:

    npm install -g nodemon
    npm install
  
## Run
### Development
(will watch server and client files, not sure if the reload is perfect for server)

    npm test

### Prod

    npm start
    
## Next steps
* Compile and run code
* Actually deal with concurrent access to sockets
