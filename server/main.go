package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("listen error: %s", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error:%s", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	var b = make([]byte, 1024)
	for {
		n, err := conn.Read(b)
		if err != nil {
			log.Printf("read error: %s", err)
			break
		} else {
			fmt.Println(string(b[:n]))
			conn.Write([]byte("+pong\r\n"))
		}
	}

	conn.Close()
}
