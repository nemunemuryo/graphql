package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  url := "https://developer.github.com/v4/explorer/"

  resp, err := http.Get(url)

  if err != nil {
    fmt.Println(err)
    return
  }

  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
}
