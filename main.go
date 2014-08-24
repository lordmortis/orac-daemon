package main

import (
  "code.google.com/p/gorest"
  "net/http"
  "github.com/lordmortis/jarvisdaemon/bridges"
)

func main() {
  gorest.RegisterService(new(bridges.BridgeService))
  http.Handle("/", gorest.Handle())
  http.ListenAndServe(":8787", nil)
}
