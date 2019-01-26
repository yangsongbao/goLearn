package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	listenr, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listenr.Accept()

		if err != nil {
			fmt.Fprintf(os.Stdout, "you got something wrong %v", err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string  // send only channel

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster()  {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <- messages:
			for cli := range clients{
				cli <- msg
			}
		case cli := <- entering:
			clients[cli] = true
		case cli := <- leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	ch <- "you are " + who

	messages <- who + "has join us"

	entering <- ch

	input := bufio.NewScanner(conn)

	for {
		for input.Scan(){
			messages <- who + ": " + input.Text()
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
