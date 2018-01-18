package main

import (
  "fmt"
  "net/http"

  "github.com/antonholmquist/jason"
)

func main() {
  //url := "http://weather.livedoor.com/forecast/webservice/json/v1?city=130020"
  url := "https://api.github.com/graphql"

  resp, err := http.Get(url)

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  v, err := jason.NewObjectFromReader(resp.Body)

  if err != nil {
        panic(err)
  }

  fmt.Print(v, "\n")

  message, err:= v.GetString("documentation_url")
  fmt.Print(message)
  //fmt.Print(v)

}
