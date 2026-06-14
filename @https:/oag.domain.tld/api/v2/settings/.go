package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  applicationId := "YOUR_applicationId_PARAMETER";
  yourOktaDomain := "subdomain.okta.com";
  reqUrl := "https://{oaghostname}/api/v2/apps/" + applicationId + "/health-check"
  req, err := http.NewRequest("GET", reqUrl, nil)
  if err != nil {
    panic(err)
  }
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }

  fmt.Println(res)
  fmt.Println(string(body))
}