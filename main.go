package main

import (
	"flag"
	"log"

	"github.com/ViktorShv95/read-adviser-bot/clients/telegram"
)

const (
	toBotHost = "api.telegram.org"
)

func main() {
	tgClient = telegram.New(toBotHost, mustToken())
}

func mustToken() string {
	token := flag.String(
		"token-bot-token", 
		"", 
		"token for access to telegram bot",
	)
	flag.Parse()
	
	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
} 