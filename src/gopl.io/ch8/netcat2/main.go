package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	go mustCopy(conn, os.Stdin)
	for {
		time.Sleep(1 * time.Minute)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalf("Copy error: %v", err)
	}
}
