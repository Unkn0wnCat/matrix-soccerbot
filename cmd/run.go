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

package cmd

import (
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/bot"

	"github.com/spf13/cobra"
)

// runCmd is the command for running the bot
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the bot",
	Long: `This runs the bot with parameters from the config file.

The bot will log in to the homeserver and start posting updates to subscribed channels.`,
	Run: func(cmd *cobra.Command, args []string) {
		bot.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
