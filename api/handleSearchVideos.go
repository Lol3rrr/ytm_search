package api

import (
  "strconv"
  "net/http"
  "encoding/json"

  "ytm_search/search"
)

func HandleSearchVideos(w http.ResponseWriter, r *http.Request) {
  keys, ok := r.URL.Query()["search"]
  if !ok || len(keys[0]) < 1 {
    w.WriteHeader(400)

    return
  }

  data, err := search.SearchVideos(keys[0])
  if err != nil {
    w.WriteHeader(400)

    return;
  }

  content, err := json.Marshal(data)
  if err != nil {
    w.WriteHeader(400)

    return;
  }

  w.Header().Set("Content-Length", strconv.Itoa(len(content)))
  w.Write(content)
}
