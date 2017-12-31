package main

import "fmt"
import "net"

func main() {
	target_host := "www.google.com"
	target_port := "80"

	fmt.Printf("acccess %v:%v", target_host, target_port)

	client, err := net.Dial("tcp", target_host+":"+target_port)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	sendMsg := "GET / HTTP/1.1\r\nHost: google.com\r\n\r\n"
	fmt.Fprintf(client, sendMsg)
	buf := make([]byte, 4096)
	n, err := client.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buf[:n]))

}
