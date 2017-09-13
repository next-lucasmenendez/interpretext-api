package main

import (
	"github.com/beeva-labs/interpretext-api/api"
	"strconv"
	"os"

	f "github.com/lucasmenendez/framework.go"
)

func main() {
	s := f.New()

	s.DebugMode(true)

	port_raw := os.Getenv("PORT")
	if port, err := strconv.Atoi(port_raw); err != nil {
		port = 9999
	} else {
		s.SetPort(port)
	}

	s.POST("/map", api.MapHandler)
	s.POST("/tweet", api.TweetHandler)
	s.POST("/summary", api.SummaryHandler)
	s.Run()
}