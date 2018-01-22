package main

import (
  "fmt"
  "net/http"

  "github.com/antonholmquist/jason"
)

func main() {
  url := "https://api.github.com/graphql"

  req, err := http.NewRequest("POST", url, nil)
  req.Header.Set("Authorization", "bearer <your token>")

  if err != nil {
    panic(err)
  }
  client := new(http.Client)
  resp, err := client.Do(req)

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  v, err := jason.NewObjectFromReader(resp.Body)

  fmt.Print(v, "\n\n")

  if err != nil {
    panic(err)
  }

  name, err:= v.GetString("repository")
  fmt.Print(name)

}
