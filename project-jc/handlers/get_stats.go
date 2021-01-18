package handlers

import (
  "log"
  "net/http"
  "encoding/json"
  "public-go/project-jc/data"
)

/*
  A GET to /stats should accept no data; it should return a JSON data structure
  for the total hash requests since server start and the average time of a hash
  request in milliseconds.
*/

type StatsResponse struct {
  Total int         `json:"total"`
  Average int64     `json:"average"`
}

// Returns a createe and populated JSON StatsResponse
func getStatsResponse(d *data.Data) []byte {

  t, a := d.GetStats()

  resp := &StatsResponse{Total:t, Average:a}
  jResp, err := json.Marshal(resp)
  if err != nil {
    log.Println("error in Marshal of StatsResponse:", err)
  }
  return jResp
}

// Handler for the /stats endpoint
func GetStats(d *data.Data) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // Write the response, the JSON StatsResponse struct
    w.WriteHeader(http.StatusOK)
    w.Write(getStatsResponse(d))
  })
}
