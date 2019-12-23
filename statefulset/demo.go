package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type modeConfig struct {
	MasterMode bool `json:"masterMode"`
}

var masterMode bool

func serverMode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var mode string
	if masterMode {
		mode = "master"
	} else {
		mode = "slave"
	}
	fmt.Fprintf(w, "hello, %s!\n", mode)
}

func masterModeOnly(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello, only for master mode\n")
}

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("read config file error : ", err)
		os.Exit(-1)
	}
	mode := new(modeConfig)
	err = json.Unmarshal(data, mode)
	if err != nil {
		fmt.Println("parse json file error : ", err)
		os.Exit(-1)
	}
	masterMode = mode.MasterMode
	fmt.Println(masterMode)
	router := httprouter.New()
	router.GET("/api/v1/mode", serverMode)
	if masterMode {
		router.GET("/api/v1/master-mode", masterModeOnly)
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
