package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ogier/pflag"

	"./action"
	"./g13"
)

var initialConfig = pflag.StringP("toml", "t", "default", "Default config to start with")

func main() {
	pflag.Parse()

	eventHandler := action.NewHandler()
	defer eventHandler.Close()

	config, err := LoadTOMLConfig(*initialConfig)
	if err != nil {
		log.Println(err)
		return
	}

	if errs, ok := config.Validate(); !ok {
		for _, err := range errs {
			log.Println(err)
			return
		}
	}

	config.ToActions(eventHandler, eventHandler.Actions)

	g13.FindDevices(eventHandler)
	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt)
	<-quitCh
	fmt.Println("")
}
