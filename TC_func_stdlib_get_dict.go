package ThreadComputation

import (

)

func TCDictGet(x interface{}, y interface{}) interface{} {
  v := x.(*TCDict).D.Get(y.(string))
  if v == nil {
    return MakeNone()
  }
  return v
}

func init() {
  RegisterOperatorCallback("get", Dict, String, TCDictGet)
}
