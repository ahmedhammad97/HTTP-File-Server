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
  packet1, err := bufio.NewReader(conn).ReadBytes('\n')

  if err != nil {
    fmt.Println("Error in reading request .. Possible corruption")
  } else {
    // remove break char from the end
    StringifiedPacket := string(packet1)[:len(packet1)-1]
    fmt.Println(StringifiedPacket)
    if strings.Contains(StringifiedPacket, "200"){
      //MUST LOOK FOR EOF TO READ ALL FILE
      packet2 := bufio.NewReader(conn)
      for {
        line, err := packet2.ReadString('\n')
        if len(line) == 0 && err != nil {
          if err == io.EOF {
            break
          }
          fmt.Println("Error in reading request .. Possible corruption")
        }
        line = strings.TrimSuffix(line, "\n")
        fmt.Println(line)
        if err != nil {
          if err == io.EOF {
            break
          }
          fmt.Println("Error in reading request .. Possible corruption")
        }
    }      // if err2 != nil {
      //   fmt.Println("Error in reading request .. Possible corruption")
      // } else {
      //   // remove break char from the end
      //   fmt.Println(len(packet2))
      //   StringifiedPacket1 := string(packet2)[:len(packet2)-1]
      //   fmt.Println(StringifiedPacket1)
      // }
    }
  }
}

func PostRoutine(source string, conn net.Conn){

}
