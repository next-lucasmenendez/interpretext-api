package main

import (
	"fmt"
	"github.com/beeva-labs/postagger"
	"log"
	"os"
	"path/filepath"
)

func train(models []string) {
	for _, modelPath := range models {
		var (
			m string = filepath.Base(modelPath)
			o string = fmt.Sprintf("%s/%s", os.Getenv("MODELS"), m)
		)
		if model, e := postagger.Train(modelPath); e != nil {
			log.Panic(e)
		} else if e := model.Store(o); e != nil {
			log.Panic(e)
		}
	}
}
