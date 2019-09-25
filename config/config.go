package config

import (
  "os"
  "encoding/json"
)

const confFile = "config.json"

type Config struct {
  YouTubeApiToken string `json:"youtubeAPIToken"`
}

func LoadConfig() (Config, error) {
  file, _ := os.Open(confFile)
  defer file.Close()

  decoder := json.NewDecoder(file)
  config := Config{}

  err := decoder.Decode(&config)
  if err != nil {
    return Config{}, err
  }

  return config, nil
}
