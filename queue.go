package main

import (
	"fmt"
	"net"
	"bufio"
	"container/list"
	"sync"
	"time"
)

func main() {
	q := queue{queueList:list.New(), wait:time.Duration(1) * time.Minute}
	go q.Writer()
	go q.Listener()
	for {
		time.Sleep(time.Duration(60) * time.Minute)
	}
}

type queue struct {
	queueList *list.List
	lock sync.RWMutex
	wait time.Duration
}


func(q *queue) Writer(){
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		q.lock.Lock()
		q.queueList.PushBack(string(message))
		q.lock.Unlock()
		conn.Write([]byte("processed\n"))
	}
}

func (q *queue) Listener() {
	// connect to this socket
	ln, _ := net.Listen("tcp", ":8082")
	conn, _ := ln.Accept()

	//conn, _ := net.Dial("tcp", "127.0.0.1:8082")
	for {
		msg := q.queueList.Front()
		if msg != nil {
			fmt.Fprint(conn, msg.Value)
		} else {
			time.Sleep(q.wait)
		}
	}
}