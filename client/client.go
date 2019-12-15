package main

import (
	"net/http"
	"time"
)

func main() {
	for {
		http.Get("http://localhost:1323/")
		time.Sleep(1 * time.Second)
	}

}
