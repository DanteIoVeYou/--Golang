package main 

import (
  "fmt"
  "net"
)

type Server struct {
  Ip string 
  Port int
}

func NewServer(ip string, port int) *Server {
  srv := Server {
    Ip: ip,
    Port: port,
  }
  return &srv
}

func (this *Server) handler(conn net.Conn) {
  fmt.Println("get a new link...")
}

func (this *Server) Start() {
  // listen
  listen_sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
  if err != nil {
    fmt.Println("listen error: ", err)
    return
  }
  // close
  defer listen_sock.Close()
  // accept
  for {
    sock, err := listen_sock.Accept()
    if err != nil {
      fmt.Println("accept error: ", err)
      continue
    }
    // handler
    go this.handler(sock)
  }
}
