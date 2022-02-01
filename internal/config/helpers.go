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

// getRoomConfigs parses and returns the bot.rooms config-key
func getRoomConfigs() RoomConfigTree {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))
	var roomConfigs RoomConfigTree

	err := viper.UnmarshalKey("bot.rooms", &roomConfigs)
	if err != nil {
		log.Panicln(p.Sprintf("Corrupted configuration: Could not load room configurations!\n%v", err))
	}

	return roomConfigs
}

// idStringToKey encodes the given ID to base32 as yaml / viper ignores case when reading
func idStringToKey(id string) string {
	return strings.ToLower(base32.StdEncoding.EncodeToString([]byte(id)))
}

// SetRoomConfigActive updates the active state for a given room
func SetRoomConfigActive(id string, active bool) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))

	// Lock room config system to prevent race conditions
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	if _, ok := roomConfigs[idStringToKey(id)]; !ok {
		roomConfigs[idStringToKey(id)] = RoomConfig{} // Create blank config if the key does not exist
	}

	roomConfig := roomConfigs[idStringToKey(id)]

	roomConfig.Active = active

	roomConfigs[idStringToKey(id)] = roomConfig // Write updated roomConfig back to roomConfigs map

	viper.Set("bot.rooms", roomConfigs) // Save to config

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}

	// Unlock room config system
	roomConfigWg.Done()
}

// GetRoomConfig returns the RoomConfig linked to the specified ID
func GetRoomConfig(id string) RoomConfig {
	roomConfigs := getRoomConfigs()

	if roomConfig, ok := roomConfigs[idStringToKey(id)]; ok {
		return roomConfig // Return a config if it exists
	}

	AddRoomConfig(id) // Config does not exist, create it

	return GetRoomConfig(id) // Run again, as now the config must exist
}

// RoomConfigInitialUpdate updates all RoomConfig entries to set activity and create blank configs
func RoomConfigInitialUpdate(ids []id.RoomID) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))

	// Lock room config system to prevent race conditions
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	// Set all active states to "false" for a blank start
	for encodedId, roomConfig := range roomConfigs {
		roomConfig.Active = false
		roomConfigs[encodedId] = roomConfig
	}

	// Go over all joined rooms
	for _, roomID := range ids {
		roomConfig, ok := roomConfigs[idStringToKey(roomID.String())]

		if !ok {
			roomConfig = RoomConfig{} // If this room has no config, create it
		}

		roomConfig.Active = true                                 // Set active state to "true"
		roomConfigs[idStringToKey(roomID.String())] = roomConfig // Save back to map
	}

	viper.Set("bot.rooms", roomConfigs)

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}

	// Unlock room config system
	roomConfigWg.Done()
}

func AddRoomConfig(id string) {
	p := message.NewPrinter(language.MustParse(viper.GetString("language")))

	// Lock room config system to prevent race conditions
	roomConfigWg.Wait()
	roomConfigWg.Add(1)

	roomConfigs := getRoomConfigs()

	if _, ok := roomConfigs[idStringToKey(id)]; ok {
		roomConfigWg.Done() // Unlock room config system
		return              // Config already exists
	}

	roomConfigs[idStringToKey(id)] = RoomConfig{Active: true}

	viper.Set("bot.rooms", roomConfigs)

	err := viper.WriteConfig()
	if err != nil {
		log.Panicln(p.Sprintf("Configuration error: Could not save configuration!"))
	}

	// Unlock room config system
	roomConfigWg.Done()
}
