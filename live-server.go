// Copyright 2015 aikah
// License MIT

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-fsnotify/fsnotify"
)

func main() {
	watchr, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	stopChannel := make(chan os.Signal, 1)

	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("watching directory %s \n", wd)
	err = watchr.Add(wd)
	if err != nil {
		log.Fatal(err)
	}
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case ev := <-watchr.Events:
			log.Println("event: ", ev)
		case err := <-watchr.Errors:
			log.Println("error: ", err)
		case <-stopChannel:
			log.Println("Exiting...")
			watchr.Close()
			os.Exit(1)
		}
	}
}
