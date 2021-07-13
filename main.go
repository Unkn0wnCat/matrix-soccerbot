package main

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,de

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"log"    
	"golang.org/x/text/language"
    "golang.org/x/text/message"
)

type formattedMessage struct {
	Type string	`json:"msgtype"`
	Body string	`json:"body"`
	Format string	`json:"format"`
	FormattedBody string	`json:"formatted_body"`
}

func init() {
	log.Print("matrix-soccerbot is initializing...")

}

func main() {
	configPtr, err := GetConfig()

	if err != nil {
		log.Fatal(err)
	}

	config := *configPtr

	lang := language.MustParse(config.Language)

	p := message.NewPrinter(lang)

	p.Printf("matrix-soccerbot has started.")

	client, err := mautrix.NewClient(config.Bot.Homeserver, id.NewUserID(config.Bot.Username, config.Bot.Homeserver), config.Bot.AccessKey)

	if err != nil {
		log.Fatal(err)
	}

	go client.Sync()


	match := GetMatchByID(61933)

	message := generateMessageForMatch(config, match)

	html := renderMarkdown(message)

	log.Println(html)

	_, err = client.SendMessageEvent(id.RoomID(config.Bot.RoomID), event.EventMessage, formattedMessage{
		"m.notice",
		message,
		"org.matrix.custom.html",
		html})

	if err != nil {
		log.Fatal(err)
	}
}