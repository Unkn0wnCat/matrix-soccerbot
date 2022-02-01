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

// RoomConfigTree is a map from string to RoomConfig
type RoomConfigTree map[string]RoomConfig

// RoomConfig is the configuration attached to every joined room
type RoomConfig struct {
	// SubscribedLeagues are leagues for which new matches should be auto-posted
	SubscribedLeagues []string `yaml:"leagues"`

	// Active tells if the bot is active in this room (Set to false on leave/kick/ban)
	Active bool `yaml:"active"`

	// Language for the bot messages in this room, this has to exist
	Language string `yaml:"lang"`
}
