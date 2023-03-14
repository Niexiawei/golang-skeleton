package main

import (
	"github.com/Niexiawei/golang-skeleton/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
