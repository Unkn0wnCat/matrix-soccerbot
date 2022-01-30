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
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

func isInRoom(client *mautrix.Client, id id.RoomID) (bool, error) {
	res, err := client.JoinedRooms()
	if err != nil {
		return false, err
	}

	for _, joinedRoom := range res.JoinedRooms {
		if joinedRoom == id {
			return true, nil
		}
	}

	return false, nil
}

func doAcceptInvite(client *mautrix.Client, id id.RoomID) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))

	roomAlreadyJoined, err := isInRoom(client, id)
	if err != nil {
		log.Println(p.Sprintf("Could not accept invite to %s due to internal error", id))
		log.Println(err)
		return
	}

	if roomAlreadyJoined {
		return
	} // If the room was already joined, ignore

	_, err = client.JoinRoom(id.String(), "", nil)
	if err != nil {
		log.Println(p.Sprintf("Could not accept invite to %s due to join error", id))
		log.Println(err)
		return
	}

	log.Println(p.Sprintf("Successfully joined room %s", id))

	config.AddRoomConfig(id.String())
}
