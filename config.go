package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type RoomConfig struct {
	RoomID            string   `yaml:"id"`
	SubscribedLeagues []string `yaml:"leagues"`
}

type Config struct {
	Language string `yaml:"language"`
	Bot      struct {
		Homeserver string       `yaml:"homeserver"`
		Username   string       `yaml:"username"`
		AccessKey  string       `yaml:"accessKey"`
		RoomID     string       `yaml:"roomID"`
		Password   string       `yaml:"password"`
		Rooms      []RoomConfig `yaml:"rooms"`
	} `yaml:"bot"`
}

func SaveConfig(config Config) error {
	file, err := os.OpenFile("config.yaml", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	content, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() (*Config, error) {
	config := &Config{}

	config.Language = "en"

	var file *os.File

	file, err := os.OpenFile("config.yaml", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		log.Println("WARNING - The config file could not be read, generating new config.")
		//return nil, err
	}

	file.Truncate(0)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	content, err := yaml.Marshal(&config)
	if err != nil {
		return nil, err
	}

	_, err = file.Write(content)
	if err != nil {
		return nil, err
	}

	return config, nil
}
