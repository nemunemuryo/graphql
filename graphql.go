package main

import (
  "fmt"
  "net/http"
  "net/url"
  "strings"

  "github.com/antonholmquist/jason"
)

func main() {
  URL := "https://api.github.com/graphql"
  values := url.Values{}
  //values := url.Values{`{repository(owner: "octocat", name: "Hello-World") {issues(last: 20, states: CLOSED) {edges {node {title url labels(first: 5) { edges { node { name}}}}}}}}`}

  // v3のapiを叩かれるっぽいから違うと思われる -----------------------------
  // values := url.Values{}
  // values.Add("user", "nemunemuryo")
  // -----------------------------------------------------------------------


  body := strings.NewReader(values.Encode())

  // 第3引数がbodyらしいからここだと思う...
  req, err := http.NewRequest("POST", URL, body)
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

  fmt.Print(v, "\n")

  if err != nil {
    panic(err)
  }

  errors, err:= v.GetObjectArray("errors")
  fmt.Print(errors)

}
