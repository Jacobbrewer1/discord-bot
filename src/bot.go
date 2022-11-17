package main

import (
	"github.com/jacobbrewer1/discord-bot/src/config"
	"log"
)

func init() {
	log.Println("initializing logging")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("logging initialized")
}

func init() {
	log.Println("reading configs")

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	log.Println("configs read")
}

func main() {

}
