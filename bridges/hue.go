package bridges

import (
  "code.google.com/p/gorest"
)

type HueBridge struct {
  Name string
  Path string
}

var bridgeList []HueBridge

type HueBridgeService struct {
  gorest.RestService `root:"/bridges/hue/" consumes:"application/json" produces:"application/json"`

  list gorest.EndPoint `method:"GET" path:"/" output:"[]HueBridge"`
}

func(serv HueBridgeService) List()[]HueBridge{
  return bridgeList
}
