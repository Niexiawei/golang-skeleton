package main

import (
	"github.com/Niexiawei/golang-skeleton/cmd"
	"github.com/Niexiawei/golang-skeleton/config"
	"github.com/Niexiawei/golang-skeleton/logger"
	"log"
)

func bootstrap() {
	config.LoadConfig()
	logger.SetupLogger()
}

func main() {
	bootstrap()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
