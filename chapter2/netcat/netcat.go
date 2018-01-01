package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "io/ioutil"
import "net"
import "os"
import "os/exec"
import "strconv"
import "strings"

//global
var upload bool = false
var (
	port              = flag.Int("p", 0, "--port")
	listen            = flag.Bool("l", false, "--listen- listen on [host]:[port]")
	command           = flag.Bool("c", false, "--command")
	uploadDestination = flag.String("u", "", "--upload=destination")
	execute           = flag.String("e", "", "--exrcute=file_to_run")
	target            = flag.String("t", "0.0.0.0", "--target ip")
)

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		usage()
	}
	//	fmt.Printf("port:%v,target:%v", *port, *target)
	if (!*listen) && (*target != "") && (*port > 0) {
		stdin := bufio.NewScanner(os.Stdin)
		buffer := ""
		//this will block, so send CTRL-D if not sending input
		for stdin.Scan() {
			buffer += stdin.Text()
		}
		clientSender(buffer)
	}
	if *listen {
		serverLoop()
	}

}

func runCommand(command string) []byte {
	command = strings.TrimRight(command, "\n")
	commands := strings.Split(command, " ")
	out, err := exec.Command(commands[0], commands[1:]...).Output()
	if err != nil {
		out = []byte(err.Error())
	}
	return out
}

func clientSender(sendMsg string) {
	client, err := net.Dial("tcp", *target+":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Printf("err :%v", err)
	}
	defer client.Close()
	if sendMsg != "" {
		client.Write([]byte(sendMsg))
	}
	for {
		response := ""
		for {
			buf := make([]byte, 4096)
			n, err := client.Read(buf)
			if err != nil {
				if err != io.EOF {
					break
				}
				fmt.Print(err)
			}
			response += string(buf[:n])
			if n < 4096 {
				break
			}
		}
		fmt.Print(response)
		//追加データ
		var c string
		_, err = fmt.Scan(&c)
		if err != nil {
			fmt.Print(err)
		}
		c += "\n"
		client.Write([]byte(c))
	}
}

func serverLoop() {
	host := *target + ":" + strconv.Itoa(*port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		fmt.Printf("ReasolveTCP %v\n", err)
		os.Exit(1)
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("ListenTCP %v\n", err)
		os.Exit(1)
	}
	defer ln.Close()

	for {
		con, err := ln.AcceptTCP()
		if err != nil {
			fmt.Printf("AcceptTCP %v\n", err)
			continue
		}
		defer con.Close()
		go clientHandler(con)
	}

}

func clientHandler(conn net.Conn) {
	if *uploadDestination != "" {
		fileBuffer := ""
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Print(err)
				break
			}
			fileBuffer += string(buf[:n])
		}
		err := ioutil.WriteFile(*uploadDestination, []byte(fileBuffer), 0664)
		if err != nil {
			conn.Write([]byte("failed saved file" + *uploadDestination + "\n"))
			fmt.Print(err)
		} else {
			conn.Write([]byte("success saved file" + *uploadDestination + "\n"))
		}
	}

	if *execute != "" {
		output := runCommand(*execute)
		conn.Write([]byte(output))
	}

	if *command {
		for {
			prompt := "<BHP:#> "
			conn.Write([]byte(prompt))
			cmdBuffer := ""
			for {
				if strings.Index(cmdBuffer, "\n") != -1 {
					break
				}
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("err :%v\n", err)
					break
				}
				if n == 0 {
					break
				}
				cmdBuffer += string(buf[:n])
			}
			response := runCommand(cmdBuffer)
			conn.Write([]byte(response))
		}
	}
}

func usage() {
	fmt.Println("BHP Net Tool")
	fmt.Println("")
	fmt.Println("Usage: ./netcat -t target_host -p port")
	fmt.Println("-l --listen							- listen on [host]:[port]")
	fmt.Println("")
	fmt.Println("-e --exrcute=file_to_run")
	fmt.Println("-c --command")
	fmt.Println("-u --upload=destination")
	fmt.Println("Example")
	fmt.Println("./netcat -t 192.168.0.1 -p 555 -l -c")
	fmt.Println("./netcat -t 192.168.0.1 -p 555 -l -u c:\\target.exe")
	fmt.Println("echo 'ABCD' | ./netcat -t 192.168.0.1 -p 555")
	os.Exit(1)
}
