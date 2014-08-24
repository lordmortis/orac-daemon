package bridges

import (
  "code.google.com/p/gorest"
)

type BridgeService struct {
  gorest.RestService `root:"/bridges/" consumes:"application/json" produces:"application/json"`

  list gorest.EndPoint `method:"GET" path:"/" output:"string"`
}

func(serv BridgeService) List() string{
  return "Hello World"
}

/*
type BridgeService struct {
  gorest.RestService `root:"/bridges/" consumes:"application/json" produces:"application/json"`

  ListBridges gorest.EndPoint `method:"GET" path:"/" output:"[]Bridge"`
}

func(serv BridgeService) ListBridges()[]Bridge{
    serv.ResponseBuilder().CacheMaxAge(60)
    return bridgeList
}
*/

