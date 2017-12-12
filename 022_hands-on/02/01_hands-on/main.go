package main

import(
	"fmt"
	"net"
	"log"
	"io"
	"bufio"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println(err.Error())
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn){
	defer conn.Close()
	
	var method string
	var uri string

	lineNumber := 1

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		if lineNumber == 1{
			method = strings.Fields(line)[0]
			uri = strings.Fields(line)[1]
		}	

		if len(line) == 0{
			break
		}

		lineNumber++
	}

	mux(conn,method,uri)
}

func mux(conn net.Conn, method string,uri string){
	if method == "GET" && uri == "/"{
		respond(conn,"<h1>GET /</h1>")
	}else if method == "GET" && uri == "/apply"{
		respond(conn,"Imagine a form here")
	}else if method == "POST" && uri == "/apply"{
		respond(conn,"Submission successful")
	} 
}

func respond(conn net.Conn,body string){
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn, body)
}