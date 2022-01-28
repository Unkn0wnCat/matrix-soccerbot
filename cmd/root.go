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

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,de

import (
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/config"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "matrix-soccerbot",
	Short: "A bot for posting summaries of soccer matches to Matrix channels.",
	Long:  `This bot allows for subscribing specific channels on Matrix to match-data from OpenLigaDB and posts updates to the channel concerning the match.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	viper.SetDefault("language", "en")
	viper.SetDefault("bot.homeserver", "https://matrix.org")
	viper.SetDefault("bot.username", "")
	viper.SetDefault("bot.password", "")
	viper.SetDefault("bot.accessKey", "")
	viper.SetDefault("bot.roomID", "")
	viper.SetDefault("bot.rooms", []config.RoomConfig{})

	rootCmd.PersistentFlags().String("language", "en", "Default language to use for logging")
	_ = viper.BindPFlag("language", rootCmd.PersistentFlags().Lookup("language"))

	cobra.OnInitialize(loadConfig)
}

func loadConfig() {
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
