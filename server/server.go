package main

import (
  "fmt"
  "bufio"
  "os"
  "net"
  //"io/ioutil"
  "time"
  "strings"
)


func main(){
  // configure port number
  port := SetPort()

  // setup TCP listener
  socket := CreateTCPListener(port)

  // kill listener when finished
  defer socket.Close()

  // wait for incomming connections
  for {
    conn, err := socket.Accept()
    if err != nil {
      fmt.Println("Incomming connection failed :(")
      time.Sleep(500 * time.Millisecond)
      continue
    } else {
      fmt.Println("Recieved connection!")
      request := PrintRequest(conn)
      if request != "" { // if request read successfully
        go HandleRequest(conn, request)
      }
    }
  }
}

func SetPort() string {
  if len(os.Args) > 1 {
    return os.Args[1]
  } else {
    return "5000"
  }
}

func CreateTCPListener(port string) net.Listener {
  for {
    // listen for tcp connections
    socket, err := net.Listen("tcp", ":" + port)
    if err == nil {
      fmt.Println("Listening to port " + port + "!...")
      return socket
    }
    // if failed to setup tcp listener, retry after 0.5 seconds
    fmt.Println("Cannot listen at port. Retrying...")
    time.Sleep(500 * time.Millisecond)
  }
}

func PrintRequest(conn net.Conn) string {
  packet, err := bufio.NewReader(conn).ReadBytes('\n')
  if err != nil {
    fmt.Println("Error in reading request .. Possible corruption")
    return ""
  } else {
    StringifiedPacket := string(packet)[:len(packet)-1]
    fmt.Println(StringifiedPacket)
    return StringifiedPacket
  }
}

func HandleRequest(conn net.Conn, req string){
  // GET / HTTP/1.0
  // POST / HTTP/1.0
  tokens := strings.Split(req, " ")
  if tokens[0] == "GET" {
    go GetRoutine(conn, tokens[1:])
  } else if tokens[0] == "POST" {
    go PostRoutine(conn, tokens[1:])
  } else {
    fmt.Println("Sorry, unsupported HTTP method!")
  }
}
