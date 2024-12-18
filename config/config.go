package config

import (
	"encoding/json"
	"os"
)

func WriteConfig(path string, data interface{}) error {
	cfg, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0655)
	if err != nil {
		return err
	}

	defer cfg.Close()

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = cfg.WriteString(string(dataBytes))
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig(name string) ([]byte, error) {
	cfg, err := os.Open(name)
	defer cfg.Close()
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil
		}

		return []byte{}, err
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
