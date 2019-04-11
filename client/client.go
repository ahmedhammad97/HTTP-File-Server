package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "net"
  "io"
)

func main (){
  reader := bufio.NewReader(os.Stdin)
  for {
    command := ReadCommand(reader)
    tokens := strings.Split(command, " ")
    socket := EstablishConnection(tokens)
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
  if len(tokens) > 3 {
    port = strings.Trim(tokens[3], " ")
  }
  host := fmt.Sprintf("%s:%s",strings.Trim(tokens[2], " "), port)
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
    // Call POST subroutine
  } else {
    panic("Unsupported command " + tokens[0])
  }
}

func GetRoutine(source string, conn net.Conn){
  conn.Write([]byte("GET /" + source + " HTTP/1.0\n"))
  message, _ := bufio.NewReader(conn).ReadString('\n')
  fmt.Println("Message recieved: " + message)
  file, err := os.Create(FixSource(source))
  if err != nil {panic(err)}
  io.Copy(file, conn)
  fmt.Println("File " + source + " stored successfully")
}

func PostRoutine(source string, conn net.Conn){

}

func FixSource(source string) string {
  return "resources/" + source
}
