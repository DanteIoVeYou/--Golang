package main 

func main() {
  srv := NewServer("127.0.0.1", 8081)
  srv.Start()
}
