package ThreadComputation

import (

)

func TCRotateSimple(v interface{}) interface{} {
  return v
}

func init() {
  RegisterFunctionCallback("rotateright", Int, TCRotateSimple)
  RegisterFunctionCallback("rotateright", Float, TCRotateSimple)
  RegisterFunctionCallback("rotateright", String, TCRotateSimple)
  RegisterFunctionCallback("rotateright", Bool, TCRotateSimple)
  RegisterFunctionCallback("rotateleft", Int, TCRotateSimple)
  RegisterFunctionCallback("rotateleft", Float, TCRotateSimple)
  RegisterFunctionCallback("rotateleft", String, TCRotateSimple)
  RegisterFunctionCallback("rotateleft", Bool, TCRotateSimple)
}
