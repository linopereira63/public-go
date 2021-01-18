package handlers

import (
  "fmt"
  "time"
  "strings"
  "net/http"
  "io/ioutil"
  "public-go/project-jc/data"
)

/*
  A POST to /hash should accept a password; it should return a job identifier
  immediate; it should then wait 5 seconds and compute the password hash. The
  hashing algorithm should be SHA512.
  Also track the time of a hash request, for use by GetStats.
*/

// Handler for the /hash endpoint
func PostHash(d *data.Data) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // Track hash time
    start := time.Now()

    defer r.Body.Close()

    val, err := ioutil.ReadAll(r.Body)
    if err != nil {
      http.Error(w, "error reading data", http.StatusBadRequest)
      return
    }

    // valildate the expected input
    sVal := strings.Split(string(val), "=")
    if len(sVal) != 2 || sVal[0] != "password" {
      http.Error(w, "invalida data", http.StatusBadRequest)
      return
    }

    // Add a job and get its jobId
    jobId := d.AddJob()

    // Compute hash in goroutine
    go d.ComputeHash(jobId, string(sVal[1]))

    // Write the response, the jobId
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%d", jobId)

    // Record the hash time
    elapsed := time.Since(start)
    d.AddToHashTime(elapsed.Nanoseconds())
  })
}
