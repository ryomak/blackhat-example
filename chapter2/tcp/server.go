package main

import "fmt"
import "net"

func main() {
	bind_ip := "0.0.0.0"
	bind_port := "1234"

	fmt.Printf("server start %v:%v\n", bind_ip, bind_port)

	server, err := net.Listen("tcp", bind_ip+":"+bind_port)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		fmt.Printf("listen from client \n")
		if err != nil {
			panic(err)
		}

		go handle_client(conn)
	}
}

func handle_client(con net.Conn) {
	buf := make([]byte, 4096)
	n, err := con.Read(buf)
	if err != nil {
		panic(err)
	}
	if n != 0 {
		fmt.Printf(string(buf[:n]))
	}
	con.Write([]byte("ban!!"))
}
