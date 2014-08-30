package bridges

import (
  "code.google.com/p/gorest"
  "strconv"
)

var bridgeTypes []BridgeType
var bridgeList []Bridge

type BridgeService struct {
  gorest.RestService `root:"/bridges/" consumes:"application/json" produces:"application/json"`

  list gorest.EndPoint `method:"GET" path:"/" output:"[]Bridge"`
  types gorest.EndPoint `method:"GET" path:"/types" output:"[]BridgeType"`  
}

func(serv BridgeService) List()[]Bridge{
  return bridgeList
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

type Bridge interface {
  id() int32
  name() string
  kind() string
}

type BridgeType struct {
  Name string
  BridgeInterface interface
}

type BridgeInterface interface {
  parameters()
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

