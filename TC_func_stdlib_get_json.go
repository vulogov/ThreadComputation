package ThreadComputation

import (

)

func TCJsonGet(x interface{}, y interface{}) interface{} {
  v := x.(*TCJson).J.Path(y.(string)).Data()
  if v == nil {
    return MakeNone()
  }
  return v
}

func init() {
  RegisterOperatorCallback("get", Json, String, TCJsonGet)
}
