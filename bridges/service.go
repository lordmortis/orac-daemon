package bridges

import (
  "code.google.com/p/gorest"
  "strconv"
)

var bridgeList []Bridge
var number int

type BridgeService struct {
  gorest.RestService `root:"/bridges/" consumes:"application/json" produces:"application/json"`

  list gorest.EndPoint `method:"GET" path:"/" output:"[]Bridge"`
  add gorest.EndPoint `method:"GET" path:"/add" output:"Bridge"`
}

func(serv BridgeService) List()[]Bridge{
  return bridgeList
}

func(serv BridgeService) Add()Bridge{
  number += 1
  var temp = Bridge{Name: "Bridge" + strconv.FormatInt(int64(number), 10), Kind: "Hue"}
  bridgeList = append(bridgeList, temp)
  return temp;
}

type Bridge struct {
  Name string
  Kind string
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

