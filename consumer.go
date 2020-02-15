package main

import (
	"net"
	"bufio"
	"fmt"
)

func main()  {

	ln, _ := net.Dial("tcp", "127.0.0.1:8081")

	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message is %s", message)
	}
	defer ln.Close()

}
