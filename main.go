package main

import (
	"github.com/kuno989/friday/cmd"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetLevel(log.DebugLevel)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}