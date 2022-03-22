package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ogier/pflag"

	"cork/go13/action"
	"cork/go13/g13"
)

func main() {
	initialConfig := pflag.StringP("toml", "t", "default", "Default config to start with")
	configFolder := pflag.StringP("toml-path", "p", "./config", "Storage folder for configuration files, defaults to working directory")
	help := pflag.BoolP("help", "h", false, "Display this message")

	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(1)
	}

	config, err := action.NewConfig(*configFolder)
	if err != nil {
		fmt.Println(err)
		return
	}

	eventHandler := config.NewHandler()
	defer eventHandler.Close()

	err = config.LoadTOMLConfig(*initialConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	if errs, ok := config.Actions.Validate(); !ok {
		for _, err := range errs {
			fmt.Println(err)
			return
		}
	}

	config.ToActions(eventHandler, eventHandler.Actions)

	g13.FindDevices(eventHandler)

	startWebServer(config)

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt)
	<-quitCh
	fmt.Println("")
}
