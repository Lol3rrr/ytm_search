package search

import (
  "errors"
)

func GetVideo(id string) (SearchVideo, error) {
  service, err := getYTService()
  if err != nil {
    return SearchVideo{}, err
  }

  call := service.Videos.List("snippet").Id(id)

  response, err := call.Do()
  if err != nil {
    return SearchVideo{}, err
  }

  if len(response.Items) < 1 {
    return SearchVideo{}, errors.New("Youtube returned empty response")
  }

  video := response.Items[0]
  snippet := video.Snippet

  result := SearchVideo{
    ID: video.Id,
    Title: snippet.Title,
    Channel: snippet.ChannelTitle,
  }

  return result, nil
}
