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
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/messageCreator"
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/openLigaDbClient"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"os"
)

func Run() {
	lang := language.MustParse(viper.GetString("language"))

	p := message.NewPrinter(lang)

	if viper.GetString("bot.homeserver") == "" || viper.GetString("bot.username") == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user identification (homeserver / username)"))
		os.Exit(1)
		return
	}

	if viper.GetString("bot.accessKey") == "" && viper.GetString("bot.password") == "" {
		log.Println(p.Sprintf("matrix-soccerbot is missing user credentials (access-key / password)"))
		log.Println(p.Sprintf("Please provide either an access-key or password"))
		os.Exit(1)
		return
	}

	log.Println(p.Sprintf("matrix-soccerbot has started."))

	client, err := mautrix.NewClient(viper.GetString("bot.homeserver"), id.NewUserID(viper.GetString("bot.username"), viper.GetString("bot.homeserver")), viper.GetString("bot.accessKey"))

	if err != nil {
		log.Println(p.Sprintf("matrix-soccerbot couldn't initialize matrix client, please check credentials"))
		log.Fatal(err)
		return
	}

	if viper.GetString("bot.accessKey") == "" {
		res, err := client.Login(&mautrix.ReqLogin{
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

		viper.Set("bot.accessKey", accessToken)
		err = viper.WriteConfig()
		if err != nil {
			log.Println(p.Sprintf("matrix-soccerbot could not save the accessKey to config"))
			log.Fatal(err)
			return
		}
	}

	go client.Sync()

	match := openLigaDbClient.GetMatchByID(61933)

	msg := messageCreator.GenerateMessageForMatch(viper.GetString("language"), match)

	html := messageCreator.RenderMarkdown(msg)

	log.Println(html)

	_, err = client.SendMessageEvent(id.RoomID(viper.GetString("bot.roomId")), event.EventMessage, formattedMessage{
		"m.notice",
		msg,
		"org.matrix.custom.html",
		html})

	if err != nil {
		log.Fatal(err)
	}
}

type formattedMessage struct {
	Type          string `json:"msgtype"`
	Body          string `json:"body"`
	Format        string `json:"format"`
	FormattedBody string `json:"formatted_body"`
}