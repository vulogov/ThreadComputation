package ThreadComputation

import (

)



func TCConvertStringToList(x interface{}) *TCList {
  switch e := x.(type) {
  case string:
    res := new(TCList)
    for i := 0; i < len(e); i++ {
      res.Add(string(e[i]))
    }
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(String, TCConvertStringToList)
}
