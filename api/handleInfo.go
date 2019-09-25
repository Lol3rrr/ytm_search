package api

import (
  "strings"
  "strconv"
  "net/http"
  "encoding/json"

  "ytm_search/info"
)

func HandleInfo(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path
  urlParts := strings.Split(url, "/")
  id := urlParts[len(urlParts) - 1]

  if id == "" || id == "get" || len(id) != 11 {
    w.WriteHeader(400)

    return
  }

  data, err := info.GetVideoInfo(id)
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
