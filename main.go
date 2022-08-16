package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var build = "develop"

func main() {
	log.Println("starting service", build)
	defer log.Println("service ended")

	showtown := make(chan os.Signal, 1)
	signal.Notify(showtown, syscall.SIGINT, syscall.SIGTERM)
	<-showtown

	log.Println("stopping service")
}
