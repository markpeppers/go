package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

var port int

func Init() {
	flag.IntVar(&port, "port", 8000, "the port")
	flag.Parse()
}

func main() {
	Init()
	fmt.Printf("Listening on port %d\n", port)

	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, phrase string, delay time.Duration) {
	io.WriteString(c, "\t"+strings.ToUpper(phrase)+"\n")
	time.Sleep(delay)
	io.WriteString(c, "\t"+phrase+"\n")
	time.Sleep(delay)
	io.WriteString(c, "\t"+strings.ToLower(phrase)+"\n")
}

func handleConn(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		echo(c, scanner.Text(), 2*time.Second)
	}
}
