package server

import (
	"fmt"
	"go-web-hook/src/task"
	"go-web-hook/src/unit"
	"net/http"
	"os"
)

func StartService(address string, port string) {
	fmt.Printf("http://%s:%s\n", address, port)
	http.HandleFunc("/", index)
	http.HandleFunc("/webhook", webHook)
	listenErr := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil)
	if listenErr != nil {
		fmt.Println(listenErr)
		os.Exit(1)
	}
}

func index(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(res, unit.GetTime())
}

func webHook(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintln(res, "success")
		go task.Add()
	} else {
		fmt.Fprintln(res, "web hook only post")
	}
}