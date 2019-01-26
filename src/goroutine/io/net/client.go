package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main()  {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connecting to ")

	done := make(chan string)

	go handleWrite(conn, done)

	go handleRead(conn, done)

}

func handleWrite(conn net.Conn, done chan string) {
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
		fmt.Println("send")
	}
}

func handleRead(conn net.Conn, done chan string) {
	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error to read message because of ", err)
			return
		}
		fmt.Println(string(buf[:reqLen-1]))
		fmt.Println("read")
	}
}
