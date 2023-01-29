package config

import (
	"encoding/json"
	"os"
)

func (cfg *Config) LoadConfig(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	stat, _ := file.Stat()

	data := make([]byte, stat.Size())

	_, err = file.Read(data)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(data), cfg)

	return err
}
