package search

import (
  "fmt"

  "google.golang.org/api/youtube/v3"

  "ytm_search/info"
)

// Retrieve playlistItems in the specified playlist
func playlistItemsList(service *youtube.Service, part string, playlistId string, pageToken string) (*youtube.PlaylistItemListResponse, error) {
  call := service.PlaylistItems.List(part)
  call = call.PlaylistId(playlistId)
  call = call.MaxResults(50)
  if pageToken != "" {
    call = call.PageToken(pageToken)
  }
  response, err := call.Do()
  if err != nil {
    return nil, err
  }
  return response, nil
}

func GetPlaylist(playlistId string) (SearchResult, error) {
  service, err := getYTService()
  if err != nil {
    return SearchResult{}, err
  }

  result := SearchResult{
    Videos: make([]SearchVideo, 0),
  }

  nextPageToken := ""
  for {
    // Retrieve next set of items in the playlist.
    playlistResponse, err := playlistItemsList(service, "snippet", playlistId, nextPageToken)
    if err != nil {
      fmt.Printf("Error: %v \n", err)

      return SearchResult{}, err
    }

    for _, playlistItem := range playlistResponse.Items {
      vidID := playlistItem.Snippet.ResourceId.VideoId

      tmpInfo, err := info.GetVideoInfo(vidID)
      if err != nil {
        continue
      }

      tmpResult := SearchVideo{
        ID: tmpInfo.ID,
        Title: tmpInfo.Title,
        Channel: tmpInfo.Channel,
      }

      result.Videos = append(result.Videos, tmpResult)
    }

    // Set the token to retrieve the next page of results
    // or exit the loop if all results have been retrieved.
    nextPageToken = playlistResponse.NextPageToken
    if nextPageToken == "" {
      break
    }
  }

  return result, nil
}
