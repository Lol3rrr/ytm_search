package search

import (
  "net/http"

  "google.golang.org/api/youtube/v3"
  "google.golang.org/api/googleapi/transport"

  "ytm_search/config"
)

type SearchVideo struct {
  ID      string `json:"id"`
  Title   string `json:"title"`
  Channel string `json:"channel"`
}

type SearchResult struct {
  Videos []SearchVideo `json:"videos"`
}

func getYTService() (*youtube.Service, error) {
  config, err := config.LoadConfig()
  if err != nil {
    return nil, err
  }

  client := &http.Client{
    Transport: &transport.APIKey{Key: config.YouTubeApiToken},
  }
  service, err := youtube.New(client)

  return service, err
}
