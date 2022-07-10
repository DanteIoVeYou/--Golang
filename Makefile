CC=go
server: main.go server.go 
	$(CC) build -o server main.go server.go 
.PHONY:clean
clean:
	rm -f server
		

