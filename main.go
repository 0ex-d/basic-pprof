package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	worker := func() {
		for i := 0; i < 1e6; i++ {
			_ = fmt.Sprintf("Number: %d", i)
		}
	}

	go func() {
		for {
			worker()
			time.Sleep(1 * time.Second)
		}
	}()
	fmt.Println("Starting pprof server at http://localhost:6060")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
