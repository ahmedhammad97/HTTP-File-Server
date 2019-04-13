#HTTP Client and server

This project is a protocol for client server file exchange

### Prerequisites

you will need to install Golang follow instructions for [Go](https://golang.org/doc/install)

### Installing

first you need to copy the project file to Go work space

then open your terminal and navigate to Project folder

```bash
$cd <folder path>
```

then run

```bash
$go build
```

you will need to run the server side

```bash
$go run server.go
```

then you need to run the client side on the client device

```bash
$go run client.go
```

now your protocol is running

## Running the tests

Enter the http v1.0 request sintax on the client side as follows

For GET request

```
GET file-name host-name (port-number)
```

For Post request

```
POST file-name host-name (port-number)
```

## Built With

- [Go](https://golang.org/doc/)

## Authors

- **Billie Thompson** - _Initial work_ - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.
