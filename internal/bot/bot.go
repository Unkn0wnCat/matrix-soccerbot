/*
 * Copyright © 2022 Kevin Kandlbinder.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bot

import (
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/config"
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/messageCreator"
	"github.com/Unkn0wnCat/matrix-soccerbot/openLigaDbClient"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Run starts the bot, blocking until an interrupt or SIGTERM is received
func Run() {
	// Set up signal channel for controlled shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Save Timestamp for filtering
	startTs := time.Now().Unix()

	// Set up l10n
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))

	checkConfig(p)

	log.Println(p.Sprintf("matrix-soccerbot has started."))

	// Initialize client, this does not check access key!
	matrixClient, err := mautrix.NewClient(
		viper.GetString("bot.homeserver"),
		id.NewUserID(viper.GetString("bot.username"), viper.GetString("bot.homeserver")),
		viper.GetString("bot.accessKey"))
	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot couldn't initialize matrix client, please check credentials"))
		log.Fatal(err)
		return
	}

	// If no accessKey is set, perform login
	if viper.GetString("bot.accessKey") == "" {
		performLogin(matrixClient, p)
	}

	// Set up sync handlers for invites and messages
	syncer := matrixClient.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, handleMessageEvent(matrixClient, startTs))
	syncer.OnEventType(event.StateMember, handleMemberEvent(matrixClient, startTs))

	// Set up async tasks
	go startSync(matrixClient, p)
	go doInitialUpdate(matrixClient, p)

	// The following is here for reference
	match, err := openLigaDbClient.GetMatchByID(61933)
	if err != nil {
		log.Fatal(err)
	}

	/*msg*/
	_ = messageCreator.GenerateMessageForMatch(viper.GetString("language"), *match)

	/*html := messageCreator.RenderMarkdown(msg)

	log.Println(html)

	_, err = matrixClient.SendMessageEvent(id.RoomID(viper.GetString("bot.roomId")), event.EventMessage, formattedMessage{
		"m.notice",
		msg,
		"org.matrix.custom.html",
		html})

	if err != nil {
		log.Fatal(err)
	}*/

	<-c
	log.Println(p.Sprintf("Shutting down..."))

	matrixClient.StopSync()

	log.Println(p.Sprintf("Goodbye!"))

	os.Exit(0)
}

// checkConfig applies constraints to the configuration and exits the program on violation
func checkConfig(p *message.Printer) {
	// Both homeserver and username are required!
	if viper.GetString("bot.homeserver") == "" || viper.GetString("bot.username") == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user identification (homeserver / username)"))
		os.Exit(1)
		return
	}

	// Either accessKey or password are required
	if viper.GetString("bot.accessKey") == "" && viper.GetString("bot.password") == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user credentials (access-key / password)"))
		log.Println(p.Sprintf("Please provide either an access-key or password"))
		os.Exit(1)
		return
	}
}

// performLogin logs in the mautrix.Client using the username and password from config
func performLogin(matrixClient *mautrix.Client, p *message.Printer) {
	res, err := matrixClient.Login(&mautrix.ReqLogin{
		Type:                     "m.login.password",
		Identifier:               mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: viper.GetString("bot.username")},
		Password:                 viper.GetString("bot.password"),
		StoreCredentials:         true,
		InitialDeviceDisplayName: "github.com/Unkn0wnCat/matrix-soccerbot",
	})
	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot couldn't sign in, please check credentials"))
		log.Fatal(err)
		return
	}

	accessToken := res.AccessToken

	// Save accessKey to configuration
	viper.Set("bot.accessKey", accessToken)
	err = viper.WriteConfig()
	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot could not save the accessKey to config"))
		log.Fatal(err)
		return
	}
}

// doInitialUpdate updates the config right after startup to catch up with joined/left rooms
func doInitialUpdate(matrixClient *mautrix.Client, p *message.Printer) {
	resp, err := matrixClient.JoinedRooms()
	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot could not read joined rooms, something is horribly wrong"))
		log.Fatalln(err)
	}

	// Hand-off list to config helper
	config.RoomConfigInitialUpdate(resp.JoinedRooms)
}

// handleMessageEvent wraps message handler taking the mautrix.Client and start timestamp as parameters
func handleMessageEvent(matrixClient *mautrix.Client, startTs int64) mautrix.EventHandler {
	return func(source mautrix.EventSource, evt *event.Event) {
		if evt.Timestamp < (startTs * 1000) {
			// Ignore old events
			return
		}

		// Cast event to correct event type
		content, ok := evt.Content.Parsed.(*event.MessageEventContent)

		if !ok {
			log.Println("Uh oh, could not typecast m.room.member event content...")
			return
		}

		username, _, err := matrixClient.UserID.Parse()
		if err != nil {
			log.Panicln("Invalid user id in client")
		}

		if !strings.HasPrefix(content.Body, "!"+username) &&
			!strings.HasPrefix(content.Body, "@"+username) &&
			!(strings.HasPrefix(content.Body, username) && strings.HasPrefix(content.FormattedBody, "<a href=\"https://matrix.to/#/"+matrixClient.UserID.String()+"\">")) {
			return
		}

		handleCommand(content.Body, evt.Sender, evt.RoomID, matrixClient)

	}
}

// handleMemberEvent wraps m.room.member (invite, join, leave, ban etc.) handler taking the mautrix.Client and start timestamp as parameters
func handleMemberEvent(matrixClient *mautrix.Client, startTs int64) func(source mautrix.EventSource, evt *event.Event) {
	return func(source mautrix.EventSource, evt *event.Event) {
		if *evt.StateKey != matrixClient.UserID.String() {
			return
		} // This does not concern us, as we are not the subject
		if evt.Timestamp < (startTs * 1000) {
			// Ignore old events, TODO: Handle missed invites.
			return
		}

		// Cast event to correct event type
		content, ok := evt.Content.Parsed.(*event.MemberEventContent)

		if !ok {
			log.Println("Uh oh, could not typecast m.room.member event content...")
			return
		}

		// If it is an invite, accept it
		if content.Membership == event.MembershipInvite {
			doAcceptInvite(matrixClient, evt.RoomID)
			return
		}

		// If it is our join event, set room to active
		if content.Membership == event.MembershipJoin {
			config.SetRoomConfigActive(evt.RoomID.String(), true)
			return
		}

		// If we left or got banned, set room to inactive
		if content.Membership.IsLeaveOrBan() {
			config.SetRoomConfigActive(evt.RoomID.String(), false)
			return
		}
	}
}

// startSync starts the mautrix.Client sync for receiving events
func startSync(matrixClient *mautrix.Client, p *message.Printer) {
	err := matrixClient.Sync()
	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot has encountered a fatal error whilst syncing"))
		log.Println(err)
		os.Exit(2)
	}
	log.Println("sync exited.")
}

// formattedMessage is a helper struct for sending HTML content
type formattedMessage struct {
	Type          string `json:"msgtype"`
	Body          string `json:"body"`
	Format        string `json:"format"`
	FormattedBody string `json:"formatted_body"`
}
