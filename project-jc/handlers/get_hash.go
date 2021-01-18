package handlers

import (
  "strings"
  "net/http"
  "strconv"
  "public-go/project-jc/data"
)

/*
  A GET to /hash should accept a job identifier; it should return the base64
  encoded password hash for the corresponding POST request.
*/

// Handler for the /hash/{jobId} endpoint
func GetHash(d *data.Data) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // Validate input
    suffix := strings.TrimPrefix(r.URL.Path, "/hash/")

    if strings.Contains(suffix, "/") {
      http.NotFound(w, r)
      return
    }

    if suffix == "" {
      http.Error(w, "missing job ID", http.StatusBadRequest)
      return
    }

    id, err := strconv.Atoi(suffix)
    if err != nil {
        http.Error(w, "invalid job ID", http.StatusBadRequest)
        return
    }

    val, present := d.Get(id)
    if !present {
      http.Error(w, "job ID not found", http.StatusBadRequest)
      return
    }

    // Write the response, the computed hash value
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(val))
  })
}
