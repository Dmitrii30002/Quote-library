package config

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type Config struct {
	DataBase struct {
		Host    string
		Port    string
		Name    string
		User    string
		Pwd     string
		Sslmode string
	}
	Server struct {
		Host string `json:"Host"`
		Port string `json:"Port"`
	} `json:"Server"`
}

func New(env_path string, json_path string) (*Config, error) {
	err := LoadEnv(env_path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DataBase: struct {
			Host    string
			Port    string
			Name    string
			User    string
			Pwd     string
			Sslmode string
		}{
			Host:    os.Getenv("DB_HOST"),
			Port:    os.Getenv("DB_PORT"),
			Name:    os.Getenv("DB_NAME"),
			User:    os.Getenv("DB_USER"),
			Pwd:     os.Getenv("DB_PASSWORD"),
			Sslmode: os.Getenv("SSLMODE"),
		},
	}

	data, err := os.ReadFile(json_path)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	return cfg, nil
}

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()
}
