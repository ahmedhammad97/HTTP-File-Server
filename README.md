# HTTP Client and server

This project is an implementation for a web file server, and a web client, that uses simple version of HTTP protocol to send and receive files.

### Prerequisites

1. You will need to install Golang, follow instructions from [Go Docs.](https://golang.org/doc/install)

2. Copy the project file to Go workspace located (usually) in your home directory.

3. From your terminal, navigate to project's folder.


### Running the Server
Head to the server folder

```bash
$cd <server>
```

Then run

```bash
$go build
```

A binary file should be created in the same directory, which then you can run using:

```bash
$./server (optional port number)
```

Now you should have the server running.

### Running the Client
From the project's folder, head to the client folder

```bash
$cd <client>
```

Then run

```bash
$go build
```

A binary file should be created in the same directory, which then you can run using:

```bash
$./client
```

Now you should have the client running.

### Running the tests

In the running client terminal, enter the HTTP/1.0 request as follows:

For GET request

```
GET file-name host-name (port-number)
```

For Post request

```
POST file-name host-name (port-number)
```



## Resources 

- [Go Docs](https://golang.org/doc/)

## Authors

- Ahmed Hammad
- Mohamed Tarek
- Ismail El-Yamany
