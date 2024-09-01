package main

import (
	"flag"
	"log"
	tgClient "telegramBot/clients/telegram"
	event_consumer "telegramBot/consumer/event-consumer"
	"telegramBot/events/telegram"
	"telegramBot/storage/files"
)

// 7056673288:AAGuxc9qhbJenoYHeQmoScFSeuexGSBkgeY

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	} 
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("bot is not specified")
	}

	return *token
}
