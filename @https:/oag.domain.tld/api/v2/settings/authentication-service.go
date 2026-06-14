package main

import (
  "fmt"
  "bytes"
  "net/http"
  "io/ioutil"
)

func main() {
  yourOktaDomain := "subdomain.okta.com";
  reqUrl := "https://{oaghostname}/api/v2/settings/authentication-service"
  var data = []byte(`{
    "serverNodes": [
      "worker-node1.test-oag.com"
    ],
    "database": {
      "type": "mysql",
      "host": "db.domain.tld",
      "port": 3306,
      "username": "<database-username>",
      "password": "<database-password>",
      "name": "oagofflinemode"
    },
    "publicDomain": "offline-idp-service.domain.tld",
    "certificateId": "15cc2bc6-b280-4d94-a0bf-c91751b40d9c"
  }`)
  req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
  if err != nil {
    panic(err)
  }
  req.Header.Add("Content-Type", "application/json")
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