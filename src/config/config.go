package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

// ReadConfig takes in an optional path to the config otherwise it will use the default
func ReadConfig(path ...string) error {

	var config = "./config/config.json"
	if path != nil {
		config = path[0]
	}

	if abs, exists := findFile(config); exists {
		log.Println("config detected at:", abs)

		c, err := os.ReadFile(abs)
		if err != nil {
			return err
		}

		log.Println(string(c))

		err = json.Unmarshal(c, &Cfg)
		if err != nil {
			return err
		}

	} else {
		return errors.New("no config found")
	}

	return nil
}

func findFile(path string) (string, bool) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}

	file, err := os.Open(abs)
	if err != nil {
		return "", false
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return abs, true
}
