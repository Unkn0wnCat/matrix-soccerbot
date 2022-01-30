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

package config

import (
	"encoding/base32"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"maunium.net/go/mautrix/id"
	"strings"
	"sync"
)

var (
	roomConfigWg sync.WaitGroup
)

func getRoomConfigs() RoomConfigTree {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))
	var roomConfigs RoomConfigTree

	err := viper.UnmarshalKey("bot.rooms", &roomConfigs)
	if err != nil {
		log.Panicln(p.Sprintf("Corrupted configuration: Could not load room configurations!\n%v", err))
	}

	return roomConfigs
}

func idStringToKey(id string) string {
	return strings.ToLower(base32.StdEncoding.EncodeToString([]byte(id)))
}

func SetRoomConfigActive(id string, active bool) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	if _, ok := roomConfigs[idStringToKey(id)]; !ok {
		roomConfigs[idStringToKey(id)] = RoomConfig{}
	}

	roomConfig := roomConfigs[idStringToKey(id)]

	roomConfig.Active = active

	roomConfigs[idStringToKey(id)] = roomConfig

	viper.Set("bot.rooms", roomConfigs)

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}
	roomConfigWg.Done()
}

func GetRoomConfig(id string) RoomConfig {
	roomConfigs := getRoomConfigs()

	if roomConfig, ok := roomConfigs[idStringToKey(id)]; ok {
		return roomConfig
	}

	AddRoomConfig(id)

	return GetRoomConfig(id) // Run again
}

func RoomConfigInitialUpdate(ids []id.RoomID) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	for encodedId, roomConfig := range roomConfigs {
		roomConfig.Active = false
		roomConfigs[encodedId] = roomConfig
	}

	for _, roomID := range ids {
		roomConfig, ok := roomConfigs[idStringToKey(roomID.String())]

		if !ok {
			roomConfig = RoomConfig{}
		}

		roomConfig.Active = true
		roomConfigs[idStringToKey(roomID.String())] = roomConfig
	}

	viper.Set("bot.rooms", roomConfigs)

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}
	roomConfigWg.Done()
}

func AddRoomConfig(id string) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	if _, ok := roomConfigs[idStringToKey(id)]; ok {
		roomConfigWg.Done()
		return // Config already exists
	}

	roomConfigs[idStringToKey(id)] = RoomConfig{Active: true}

	viper.Set("bot.rooms", roomConfigs)

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}
	roomConfigWg.Done()
}
