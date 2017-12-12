package main

import(
	"bufio"
	"fmt"
	"net"
	"log"
	"strings"
)

func main() {
	li, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for{
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}	

func handle(conn net.Conn){
	defer conn.Close()
	line := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		text := scanner.Text()
		if(line == 0){
			_ = strings.Fields(text)[0]
			path := strings.Fields(text)[1]
			fmt.Printf("Path is %s\n",path)
		}
		if(len(text) == 0){
			break
		}
		line++;
	}
}