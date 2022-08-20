package main

import (
	"fmt"
	"go.uber.org/automaxprocs/maxprocs"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		fmt.Println("maxprocs: %w", err)
		os.Exit(1)
	}

	g := runtime.GOMAXPROCS(0)
	log.Printf("starting service build[%s] CPU[%d]\n", build, g)
	defer log.Println("service ended")

	showtown := make(chan os.Signal, 1)
	signal.Notify(showtown, syscall.SIGINT, syscall.SIGTERM)
	<-showtown

	log.Println("stopping service")
}
