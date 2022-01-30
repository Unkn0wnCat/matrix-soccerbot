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
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"strings"
)

func handleCommand(command string, sender id.UserID, id id.RoomID, client *mautrix.Client) {
	username, _, err := client.UserID.Parse()
	if err != nil {
		log.Panicln("Invalid user id in client")
	}

	command = strings.TrimPrefix(command, "!")
	command = strings.TrimPrefix(command, "@")
	command = strings.TrimPrefix(command, username)
	command = strings.TrimSpace(command)

	log.Println(command)

	if strings.HasPrefix(command, "help") {
		commandHelp(sender, id, client)
		return
	}

	if strings.HasPrefix(command, "setlang") {
		commandSetLang(strings.TrimPrefix(command, "setlang"), sender, id, client)
		return
	}

	commandHelp(sender, id, client)
	return
}

func commandSetLang(params string, sender id.UserID, id id.RoomID, client *mautrix.Client) {
	// TODO
}

func commandHelp(sender id.UserID, id id.RoomID, client *mautrix.Client) {
	client.SendNotice(id, "matrix-soccerbot help\n\n!soccerbot help - shows this help")
}
