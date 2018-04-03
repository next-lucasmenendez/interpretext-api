package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string = os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No action provided. Use help for more info.")
	}

	var (
		action string   = args[0]
		params []string = args[1:]
	)

	switch action {
	case "help":
		fmt.Println("Use: ACTION [PARAMS]\n\t Available actions:\n\thelp: provide instructions for use\n\ttrain [model_paths] : trains provided models and store it on MODELS env\n\tserver : run api server")
	case "train":
		if os.Getenv("MODELS") == "" || len(params) == 0 {
			fmt.Println("Empty MODELS env variable or no models for train provided.")
		}
		train(params)
	case "server":
		startApi()
	}
}
