package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello world")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

}
