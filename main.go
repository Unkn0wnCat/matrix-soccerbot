package main

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,de

import (
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"log"    
	"golang.org/x/text/language"
    "golang.org/x/text/message"
	"os"
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
		log.Println("matrix-soccerbot can neither read nor create a configuration file")
		log.Fatal(err)
	}

	config := *configPtr

	lang := language.MustParse(config.Language)

	p := message.NewPrinter(lang)

	if config.Bot.Homeserver == "" || config.Bot.Username == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user identification (homeserver / username)"))
		os.Exit(1)
	}

	if config.Bot.AccessKey == "" && config.Bot.Password == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user credentials (access-key / password)"))
		log.Println(p.Sprintf("Please provide either an access-key or password"))
		os.Exit(1)
	}

	log.Println(p.Sprintf("matrix-soccerbot has started."))

	client, err := mautrix.NewClient(config.Bot.Homeserver, id.NewUserID(config.Bot.Username, config.Bot.Homeserver), config.Bot.AccessKey)

	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot couldn't initialize matrix client, please check credentials"))
		log.Fatal(err)
		os.Exit(1)
	}

	if config.Bot.AccessKey == "" {
		res, err := client.Login(&mautrix.ReqLogin{
			Type:             "m.login.password",
			Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: config.Bot.Username},
			Password:         config.Bot.Password,
			StoreCredentials: true,
			InitialDeviceDisplayName: "github.com/Unkn0wnCat/matrix-soccerbot",
		})

		if err != nil {
			log.Println(p.Sprintf("matrix-soccerbot couldn't sign in, please check credentials"))
			log.Fatal(err)
			os.Exit(1)
		}

		accessToken := res.AccessToken

		config.Bot.AccessKey = accessToken

		SaveConfig(config)
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