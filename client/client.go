package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "net"
  "io"
  "io/ioutil"
  "bytes"
)

func main (){
  // create reader for user input
  reader := bufio.NewReader(os.Stdin)
  for {
    command := ReadCommand(reader)
    tokens := strings.Split(command, " ")
    socket := EstablishConnection(tokens)
    // close connection when finished
    defer socket.Close()
    HandleCommand(tokens, socket)
  }
}

func ReadCommand(reader *bufio.Reader) string {
  fmt.Print("Next Command: \n")
  text, _ := reader.ReadString('\n')
  return text
}

func EstablishConnection(tokens []string) net.Conn {
  port := "80"
  // if custom port is specified
  if len(tokens) > 3 {
    port = strings.Trim(tokens[3], " ")
  }
  // sticking host. ex: 192.168.1.5:5050
  host := fmt.Sprintf("%s:%s",strings.Trim(tokens[2], " "), port)
  // create connection with server
  socket, err := net.Dial("tcp", strings.Trim(host, "\n"))
	if err != nil {
		fmt.Println("Error establishing connection:", err.Error())
	}
  return socket
}

func HandleCommand(tokens []string, conn net.Conn){
  if tokens[0] == "GET" {
    GetRoutine(strings.Trim(tokens[1], " "), conn)
  } else if tokens[0] == "POST" {
    PostRoutine(strings.Trim(tokens[1], " "), conn)
  } else {
    panic("Unsupported command " + tokens[0])
  }
}

func GetRoutine(source string, conn net.Conn){
  conn.Write([]byte("GET /" + source + " HTTP/1.0\n"))

  // read the first response (200 OK or 404 Not Found)
  packet, err := bufio.NewReader(conn).ReadBytes('\n')

  if err != nil {
    fmt.Println("Error in reading request .. Possible corruption")
    panic(err)
  }

  fmt.Println(string(packet))

  if strings.Contains(string(packet), "200"){
    file, FileErr := os.Create(FixSource(source))
    if FileErr != nil {panic(err)}
    // copy bytes comming from connection to the file
    _, CopyErr := io.Copy(file, conn)
    if CopyErr != nil {panic(CopyErr)}
  }
}

func PostRoutine(source string, conn net.Conn){
  file, err := ioutil.ReadFile(FixSource(source))
  if err != nil {
    // file was not found
    fmt.Println("File " + source + " Cannot be found")
  } else {
    // file found
    conn.Write([]byte("POST /" + source + " HTTP/1.0\n"))
    packet, err := bufio.NewReader(conn).ReadBytes('\n')
    if err != nil {
      fmt.Println("Error in reading request .. Possible corruption")
      panic(err)
    }

    fmt.Println(string(packet))

    if strings.Contains(string(packet), "200"){
      // upload file bytes to the connection
      _, err3 := io.Copy(conn, bytes.NewReader(file))
      if err3 != nil {panic(err3)}
    }
  }
}

func FixSource(source string) string {
  return "resources/" + source
}
