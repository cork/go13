package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"./action"
	"./g13"
)

func main() {
	eventHandler := action.NewHandler()
	defer eventHandler.Close()

	config, _ := ParseTOMLConfig(func(config []byte, err error) *string { script := string(config); return &script }(ioutil.ReadFile("test.toml")))

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
