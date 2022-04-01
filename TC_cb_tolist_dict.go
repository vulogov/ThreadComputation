package ThreadComputation

import (

)



func TCConvertDictToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCDict:
    res := new(TCList)
    for item := range e.D.Items() {
      p := new(TCPair)
      p.X = item.Key.(string)
      p.Y = item.Value
      res.Add(p)
    }
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(Dict, TCConvertDictToList)
}
