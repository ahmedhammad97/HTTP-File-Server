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
    conn := EstablishConnection(tokens)
    HandleCommand(tokens, conn)

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
    port = tokens[3]
  }
  conn, err := net.Dial("tcp", tokens[2]+":"+port)
	if err != nil {
		fmt.Println("Error establishing connection: ", err.Error())
	}
  return conn
}

func HandleCommand(tokens []string, conn net.Conn){
  defer conn.Close
  if tokens[0] == "GET" {
    // Call GET subroutine
  } else if tokens[0] == "POST" {
    // Call POST subroutine
  } else {
    panic("Unsupported command " + tokens[0])
    return
  }
}
