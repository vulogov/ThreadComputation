package ThreadComputation

import (

)



func TCConvertNumbersToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCNumbers:
    res := new(TCList)
    for i := 0; i < len(e.N); i++ {
      res.Add(e.N[i])
    }
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(Numbers, TCConvertNumbersToList)
}
