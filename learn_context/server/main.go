package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r * http.Request) {
	fmt.Println("handler start")
	ctx := r.Context()

	complete := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	case <- complete: // 5秒过去了
		fmt.Println("finish do something")
	case <- ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("handler end")
}


func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
