package handlers

import (
	"encoding/json"
	"log"
	"net/http"

  "github.com/tholiv/go-kube/version"
)

// home is a simple HTTP handler function which writes a response.
func home(buildTime, commit, release string) http.HanlderFunc {
  return func(w http.ResponseWriter, _ *http.Request) {
    info := struct {
      BuildTime string `json:"buildTime"`
      Commit    string `json:"commit"`
      Release   string `json:"release"`
    }{
      version.BuildTime, version.Commit, version.Release,
    }

    body, err := json.Marshal(info)
    if err != nil {
      log.Printf("Could not encode info data: %v", err)
      http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
  }
}
