// gopl Ex 8.1
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	Offset    int
	Location  string
	Address   string
	FmtString string
	C         net.Conn
}

func (c *Clock) Init() {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	c.C = conn
	fmtString := "\r"
	for i := 0; i < c.Offset; i++ {
		fmtString += "\t\t"
	}
	fmtString += "%s"
	c.FmtString = fmtString
}

func (c *Clock) Run() {
	defer c.C.Close()
	p := make([]byte, 9)
	for {
		bytes, err := c.C.Read(p)
		if bytes > 0 {
			timeStr := string(p[:8])
			fmt.Printf(c.FmtString, timeStr)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Some error occurred: %v", err)
		}
	}
}

func main() {
	fmt.Printf("\n\n\n\n")
	clocks := make([]*Clock, 0)
	for i, arg := range os.Args[1:] {
		eqPos := strings.Index(arg, "=")
		if eqPos < 0 {
			kaboom(nil)
		}
		location := arg[:eqPos]
		address := arg[eqPos+1:]
		clock := &Clock{
			Offset:   i,
			Location: location,
			Address:  address,
		}
		clock.Init()
		clocks = append(clocks, clock)
		fmt.Printf(clock.FmtString, clock.Location)
	}
	fmt.Printf("\n")
	for _, clock := range clocks {
		go clock.Run()
	}
	for {
		time.Sleep(time.Second)
	}
}

func kaboom(err error) {
	log.Fatalf("Error, stupid: %v", err)
}
