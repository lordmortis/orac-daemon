package bridges

import (
  "code.google.com/p/gorest"
)

type BridgeType struct {
  Name string
  Path string
}

var bridgeTypes []BridgeType
//var bridgeList []Bridge

func init() {
//  gorest.RegisterService(new(BridgeService))
/*  var temp = BridgeType{Name: "Hue", Path: "hue"}
  bridgeTypes = bridgeTypes.append(bridgeTypes, temp)*/
}

type BridgeService struct {
  gorest.RestService `root:"/bridges/" consumes:"application/json" produces:"application/json"`

//  list gorest.EndPoint `method:"GET" path:"/" output:"[]Bridge"`
//  types gorest.EndPoint `method:"GET" path:"/types" output:"[]BridgeType"`  
}

/*func(serv BridgeService) List()[]Bridge{
  return bridgeList
}

func(serv BridgeService) Types()[]Bridge{
  return bridgeTypes
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


type BridgeInterface interface {
  parameters()
}*/

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

