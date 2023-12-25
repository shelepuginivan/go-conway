package main

import (
	"log"

	"github.com/shelepuginivan/go-conway/cmd"
)

func main() {
	if err := cmd.Game(); err != nil {
		log.Fatal(err)
	}
}
