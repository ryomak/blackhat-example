package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	conn, _ := net.Dial("tcp", "localhost:8000")
	for {
		fmt.Fprintf(conn, "pwd"+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(strings.TrimRight(message, "\n") + ">")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		s := bufio.NewScanner(conn)
		for s.Scan() {
			fmt.Println(s.Text())
		}
	}
}
