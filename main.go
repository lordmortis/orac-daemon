package main

import (
  "code.google.com/p/gorest"
  "net/http"
  "github.com/lordmortis/orac-daemon/bridges"
)

func main() {
  gorest.RegisterService(new(bridges.BridgeService))
  gorest.RegisterService(new(bridges.HueBridgeService))
  http.Handle("/", gorest.Handle())
  http.ListenAndServe(":8787", nil)
}
