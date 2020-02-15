package main

import (
	"net"
	"fmt"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	for i := 0; i < 50; i ++ {
		conn.Write([]byte(fmt.Sprintf(`{"message": "message %d"}`, i)))
	}
	defer conn.Close()
}
