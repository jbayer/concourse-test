package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:"+os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test")
	}))
}
