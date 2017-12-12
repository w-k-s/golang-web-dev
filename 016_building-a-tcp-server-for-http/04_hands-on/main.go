package main

import (
	"bufio"
	"fmt"
	"net"
	"log"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer listener.Close()

	for{
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	multiplex(request(conn))
}

func request(conn net.Conn) (net.Conn,string,string){

	line := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		text := scanner.Text()
		if line == 0 {
			method := strings.Fields(text)[0]
			path := strings.Fields(text)[1]

			return conn,method,path
		}
	}

	return conn,"",""
}

func multiplex(conn net.Conn, method string, path string){

	if method == "GET" && path == "/home" {
		writeHTML(conn,"<h1>Welcome Home!</h1>")
	}
	if method == "GET" && path == "/saucepan"{
		writeHTML(conn, `<img src="https://i.redd.it/yptrwtadfo001.jpg"/>`)
	}
}

func writeHTML(conn net.Conn,html string) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>`+html+`</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}