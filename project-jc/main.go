package main

import (
  "log"
  "net/http"
  "time"
  "context"
  "os"
  "os/signal"
  "syscall"

  "public-go/project-jc/handlers"
  "public-go/project-jc/data"
)


type serverHandler struct {
  mux *http.ServeMux
}

// ServeHTTP to satisfy the Handler interface
func (s *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  s.mux.ServeHTTP(w, r)
}

// Creates a serverHandler, adds its enpoints and handlers, then creates
// the server configured for port 8080 and returns it
func createServer() *http.Server {

  s := &serverHandler{mux: http.NewServeMux()}

  s.mux.Handle("/hash", handlers.PostHash(d))
  s.mux.Handle("/hash/", handlers.GetHash(d))
  s.mux.Handle("/stats", handlers.GetStats(d))
  s.mux.Handle("/shutdown", shutdown())

  hs := &http.Server{Addr: ":8080", Handler: s}
  return hs
}


var d = data.NewData()
var stop = make(chan os.Signal, 1)

/*
  Creates an http server, launches its ListenAndServe via a go routine, then
  waits on the stop channel for a shutdown request or a cmd-line crtl-C
  interrupt.
*/

func main() {
  hs := createServer()

  go func() {
    log.Printf("Listening on http://0.0.0.0%s\n", hs.Addr)

    if err := hs.ListenAndServe(); err != http.ErrServerClosed {
      log.Fatal(err)
    }
  }()

  stopGracefully(hs, 10*time.Second)
}

// Stops incoming server requests, waits until all ongoing
// hash computations are completed, then exits
func stopGracefully(hs *http.Server, timeout time.Duration) {

  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  <-stop

  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()

  log.Printf("Stopping server with a %s timeout\n", timeout)

  if err := hs.Shutdown(ctx); err != nil {
    log.Printf("Error: %v\n", err)
  } else {
    // Wait until all hashes complete
    for d.IsBusy() {
      time.Sleep(500 *time.Millisecond)
    }
    log.Println("Server stopped")
  }
}

// Handler for the /shutdown endpoint, to support a graceful shutdown request
func shutdown() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Trigger a graceful shutdown
    stop <- os.Interrupt
  })
}
