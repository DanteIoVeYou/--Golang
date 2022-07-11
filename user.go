package main

import (
  "net"
)

type User struct {
  Name string
  Addr string
  C chan string
  Conn net.Conn
}

func NewUser(conn net.Conn) *User {
  userAddr := conn.RemoteAddr().String()
  user := &User {
    Name: userAddr,
    Addr: userAddr,
    C: make(chan string),
    Conn: conn,
  }
  // listen user's channel
  go user.ListenMsg()

  return user
}

func (this *User)ListenMsg() {
  for {
    msg := <- this.C
    this.Conn.Write([]byte(msg + "\n"))
  }
}

