package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"

	"os"

	"log"

	"strings"

	crypt "github.com/amoghe/go-crypt"
)

var (
	passfile   string
	dictionary string
)

func init() {
	flag.StringVar(&passfile, "f", "", "Open shadow")
	flag.StringVar(&dictionary, "d", "", "Open pass dictionary")
}

func main() {
	flag.Parse()
	if passfile == "" || dictionary == "" {
		println("Please " + os.Args[0] + " -h")
		os.Exit(0)
	}

	passFile, err := os.Open(passfile)
	if err != nil {
		log.Fatalln(err)
	}
	defer passFile.Close()

	dictFile, err := ioutil.ReadFile(dictionary)
	if err != nil {
		log.Fatalln(err)
	}

	passDict := strings.Split(string(dictFile), "\n")

	scanner := bufio.NewScanner(passFile)
	for scanner.Scan() {
		j := scanner.Text()
		if strings.Contains(j, ":") {
			user := strings.Split(j, ":")[0]
			cryptPass := strings.Split(j, ":")[1]
			fmt.Printf("[*] Cracking Password For: %v\n", user)
			for i := 0; i < len(passDict)-1; i++ {
				if testPass(cryptPass, passDict[i]) != "" {
					println(testPass(cryptPass, passDict[i]))
					break
				}
			}
		}
	}
}

func testPass(cryptPass string, passWord string) string {
	saltSearch := strings.LastIndex(cryptPass, "$")
	salt := cryptPass[0:saltSearch]

	cryptWord, err := crypt.Crypt(passWord, salt)
	if err != nil {
		log.Fatalf("Ошибка SHA: %v", err)
	}
	if cryptWord == cryptPass {
		return "[+] Found PASSWORD: " + passWord
	}
	return ""
}
