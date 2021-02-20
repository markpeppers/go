package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"strings"
)

var port int

func Init() {
	flag.IntVar(&port, "port", 8000, "the port")
	flag.Parse()
}

func main() {
	Init()
	fmt.Printf("Listening on port %d\n", port)

	/*
		// home dir contents
		root := "/Users/markpeppers"
		files, err := ioutil.ReadDir(root)
		for _, file := range files {
			fmt.Println(file.Name())
		}
	*/

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

func handleConn(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	cwd := "/Users/markpeppers"
	io.WriteString(c, "> ")
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		switch tokens[0] {
		case "ls":
			ls(c, cwd)
		case "pwd":
			pwd(c, cwd)
		case "cd":
			cd(c, &cwd, tokens[1])
		case "close":
			return
		default:
			noCmd(c, tokens[0])
		}
		io.WriteString(c, "> ")
	}
}

func ls(c net.Conn, cwd string) {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		io.WriteString(c, fmt.Sprintf("Error: %v\n", err))
		return
	}
	for _, file := range files {
		io.WriteString(c, fmt.Sprintf("%s", file.Name()))
		if file.IsDir() {
			io.WriteString(c, "/")
		}
		io.WriteString(c, "\n")
	}
}

func pwd(c net.Conn, cwd string) {
	io.WriteString(c, fmt.Sprintf("%s\n", cwd))
}

func noCmd(c net.Conn, cmd string) {
	io.WriteString(c, fmt.Sprintf("%s not a valid command\n", cmd))
}

func cd(c net.Conn, cwd *string, dir string) {
	if dir == ".." {
		pathList := strings.Split(*cwd, "/")
		*cwd = filepath.Join(pathList[:len(pathList)-1]...)
		*cwd = fmt.Sprintf("/%s", *cwd)
		io.WriteString(c, "ok\n")
		return
	}
	files, err := ioutil.ReadDir(*cwd)
	if err != nil {
		io.WriteString(c, fmt.Sprintf("Can't read the current dir: %v\n", err))
		return
	}
	for _, file := range files {
		if file.Name() == dir {
			if !file.IsDir() {
				io.WriteString(c, fmt.Sprintf("%s not a directory\n", file.Name()))
				return
			}
			*cwd = filepath.Join(*cwd, dir)
			io.WriteString(c, "ok\n")
			return
		}
	}
	io.WriteString(c, fmt.Sprintf("%s not in the current directory\n", dir))
}
