package main

import (
  "fmt"
  "time"
  "strings"
  "strconv"
  "os"
  "io/ioutil"
  "net/http"
)

/*
  Crude but functional test client. Its focus is to bombard the project-jc
  server's REST API to ensure it works under pressure.
*/

func main() {

  url := "http://localhost:8080"
  hash_url := url + "/hash"
  getHash_url := url + "/hash/"
  stats_url := url + "/stats"

  var data = "password=angryMonkey"

  getUrl(stats_url)

  for i := 0; i<1010; i++ {
    go postUrl(hash_url, data)
    time.Sleep(time.Millisecond)
  }

  // Note, only getting the first 1000 hashes
  for j := 1; j<=1000; j++ {
    get_url := getHash_url + strconv.Itoa(j)
    getUrl(get_url)
    getUrl(stats_url)

    time.Sleep(1 * time.Millisecond)
  }

  get_url := getHash_url + strconv.Itoa(1000)

  getUrl(get_url)
  getUrl(stats_url)

}


func getUrl(url string) {

  fmt.Println("Getting url:", url)

  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("Error during Get:", err)
    os.Exit(1)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Error while reading body:", err)
    return
  }

  fmt.Println(string(body))
}

func postUrl(url, s string) {

  fmt.Println("Posting url:", url)

  req, err := http.NewRequest("POST", url, strings.NewReader(s))
  if err != nil {
    fmt.Println("Error during http.NewRequest:", err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  c := &http.Client{}
  resp, err := c.Do(req)
  if err != nil {
    fmt.Println("Error during http.Do:", err)
    os.Exit(1)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Error while reading body:", err)
    return
  }

  fmt.Println(string(body))
}
