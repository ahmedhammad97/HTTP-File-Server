package main

import (
  "fmt"
  "bufio"
  "os"
  "io"
  "net"
  "io/ioutil"
  "time"
  "strings"
)

const DefaultPort string = "5050"

func main(){
  // configure port number
  port := SetPort()

  // setup TCP listener
  socket := CreateTCPListener(port)

  // kill listener when finished
  defer socket.Close()

  for {
    // wait for incomming connections
    conn, err := socket.Accept()
    if err != nil {
      fmt.Println("Incomming connection failed :(")
      time.Sleep(500 * time.Millisecond)
      continue
    } else {
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
    return DefaultPort
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
    // remove break char from the end
    StringifiedPacket := string(packet)[:len(packet)-1]
    fmt.Println(StringifiedPacket)
    return StringifiedPacket
  }
}

func HandleRequest(conn net.Conn, req string){
  // GET /index.html HTTP/1.0
  // POST / HTTP/1.0
  // close connection when finished
  defer conn.Close()
  tokens := strings.Split(req, " ")
  if tokens[0] == "GET" {
    go GetRoutine(conn, FixSource(tokens[1]))
  } else if tokens[0] == "POST" {
    go PostRoutine(conn, FixSource(tokens[1]))
  } else {
    fmt.Println("Sorry, unsupported HTTP method!")
    conn.Close()
  }
}

func GetRoutine(conn net.Conn, source string){
  file, err := ioutil.ReadFile(source)
  if err != nil {
        // file was not found
        fmt.Println(strings.Join([]string{"File", source, "Cannot be found"}, " "))
        conn.Write([]byte("HTTP/1.0 404 Not Found\r\n"))
    } else {
        // file found
        conn.Write([]byte("HTTP/1.0 200 OK\r\n"))
        conn.Write(file)
    }
}

func PostRoutine(conn net.Conn, source string){
  conn.Write([]byte("HTTP/1.0 200 OK\r\n"))
	file, err := os.Create(source)
  if err != nil {
    fmt.Println("Cannot create file " + source)
    fmt.Println(err)
    return
  }
  defer file.Close()
  io.Copy(file, conn)
  fmt.Println(strings.Join([]string{"File", source, "stored successfully"}, " "))
}

func FixSource(source string) string {
  if source == "/" {
    return "resources/index.html"
  } else {
    return strings.Join([]string{"resources", source}, "")
  }
}
