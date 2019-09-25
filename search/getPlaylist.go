package search

import (
  "fmt"
  "sync"

  "google.golang.org/api/youtube/v3"

  "ytm_search/info"
)

// Retrieve playlistItems in the specified playlist
func playlistItemsList(service *youtube.Service, playlistId string, pageToken string) (*youtube.PlaylistItemListResponse, error) {
  call := service.PlaylistItems.List("snippet")
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

func loadVidInfo(waitgroup *sync.WaitGroup, vidID string, vidIndex int, vids []SearchVideo) {
  defer waitgroup.Done()
  defer func () {
    if (recover() != nil) {
      fmt.Println("Routine Paniced")
    }
  }()

  tmpInfo, err := info.GetVideoInfo(vidID)
  if err != nil {
    return
  }

  tmpResult := SearchVideo{
    ID: tmpInfo.ID,
    Title: tmpInfo.Title,
    Channel: tmpInfo.Channel,
  }

  vids[vidIndex] = tmpResult

}

func loadPage(pageToken string, plID string, service *youtube.Service) (string, []SearchVideo) {
  // Retrieve next set of items in the playlist.
  playlistResponse, err := playlistItemsList(service, plID, pageToken)
  if err != nil {
    fmt.Printf("Error: %v \n", err)

    return "", nil
  }

  vids := make([]SearchVideo, len(playlistResponse.Items))
  var waitgroup sync.WaitGroup

  for index, playlistItem := range playlistResponse.Items {
    waitgroup.Add(1)
    vidID := playlistItem.Snippet.ResourceId.VideoId

    go loadVidInfo(&waitgroup, vidID, index, vids)
  }

  waitgroup.Wait()

  // Set the token to retrieve the next page of results
  // or exit the loop if all results have been retrieved.
  return playlistResponse.NextPageToken, vids
}

func GetPlaylist(playlistId string) (SearchResult, error) {
  fmt.Println("Loading Playlist")

  service, err := getYTService()
  if err != nil {
    return SearchResult{}, err
  }

  result := SearchResult{
    Videos: make([]SearchVideo, 0),
  }

  nextPageToken := ""
  for {
    fmt.Println("Loading Page")
    nPageToken, vids := loadPage(nextPageToken, playlistId, service)

    result.Videos = append(result.Videos, vids...)

    // Set the token to retrieve the next page of results
    // or exit the loop if all results have been retrieved.
    nextPageToken = nPageToken
    if nextPageToken == "" {
      break
    }
  }

  fmt.Println("Loaded Playlist")

  return result, nil
}
