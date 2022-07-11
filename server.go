package main

import (
  "fmt"
  "net"
  "sync"
)

type Server struct {
  Ip string
  Port int
  C chan string
  OnlineMap map[string]*User
  MapLock sync.RWMutex
}

func NewServer(ip string, port int) *Server {
  srv := Server {
    Ip: ip,
    Port: port,
    C: make(chan string),
    OnlineMap: make(map[string]*User),
  }
  return &srv
}

func (this *Server)BroadcastMsg(msg string, user *User) {
  msg = "[" + user.Name + "]" + msg
  this.C <- msg
}

func (this *Server) Handler(conn net.Conn) {
  // fmt.Println("get a new link...") 
  newuser := NewUser(conn)
  username := newuser.Addr
  // 添加用户到map
  this.MapLock.Lock()
  this.OnlineMap[username] = newuser
  this.MapLock.Unlock()
  // 广播用户登录消息
  this.BroadcastMsg(" login successfully... ", newuser)
  select {}
}


func (this *Server)SendMsg() {
  for {
    msg := <- this.C
    this.MapLock.Lock()
    for _, user := range this.OnlineMap {
      user.C <- msg
    }
    this.MapLock.Unlock()
  }
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
  // create a gotroutine to listen chan whether there is msg to wirte into User's chan
  go this.SendMsg()
  // accept
  for {
    sock, err := listen_sock.Accept()
    if err != nil {
      fmt.Println("accept error: ", err)
      continue
    }
    // handler
    go this.Handler(sock)
  }
}

