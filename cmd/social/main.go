package main

import (
	"fmt"
	"log"
	"mymod/cmd/social/service"
	"net/http"
	"os"
)

func main() {
	var port string
	f := os.Args[1:]
	for i := 0; i < len(f); i++ {
		port = f[i]
	}
	fmt.Println("Server run localhost", port)

	srv := &service.Server{}
	srv.RegisterHandlers()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
