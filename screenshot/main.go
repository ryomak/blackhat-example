package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := 3564
	http.HandleFunc("/", handler)
	fmt.Printf("server start %v \n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	if r.Method == "POST" {
		PicDecode(string(data))
	}
	fmt.Printf("%v\n", r.Method)
}

func PicDecode(str string) {
	data, _ := base64.StdEncoding.DecodeString(str)

	file, err := os.Create("./tmp/" + time.Now().Format("20060102150405") + ".jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(data)
}
