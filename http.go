package main

import (
	"io"
    "io/ioutil"
    "encoding/json"
	"net/http"
	"log"
)


type Options struct {
      Path string
      Port string
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, "hello, world!\n")
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	//err := http.ListenAndServe(":8001", nil)
	
    op := &Options{Path: "./", Port: "8001"}
    //Read config file
    data, _ := ioutil.ReadFile("./config.json")
    err := json.Unmarshal(data, op)
    if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
    log.Print("Parsed options from config file: ", op)
}

