package ThreadComputation

import (
  "fmt"
)



type TCSizeFun func(interface{}) int

func TCSizeSimple(data interface{}) int {
  return 1
}

func TCStringSize(data interface{}) int {
  switch data.(type) {
  case string:
    return len(data.(string))
  }
  return 0
}

func TCListSize(data interface{}) int {
  switch data.(type) {
  case *TCList:
    return data.(*TCList).Q.Len()
  }
  return 0
}

func TCNumbersSize(data interface{}) int {
  switch data.(type) {
  case *TCNumbers:
    return data.(*TCNumbers).Len()
  }
  return 0
}

func TCMatrixSize(data interface{}) int {
  switch data.(type) {
  case *TCMatrix:
    x, y := data.(*TCMatrix).M.Dims()
    return x * y
  }
  return 0
}

func TCNeuralSize(data interface{}) int {
  switch data.(type) {
  case *TCNeural:
    return len(data.(*TCNeural).Patterns)
  }
  return 0
}

func RegisterSizeCallback(from_type int, fun TCSizeFun) {
  fname := fmt.Sprintf("size.%v", from_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetSizeCallback(x interface{}) TCSizeFun {
  var fn string
  switch x.(type) {
  case int64:
    fn = fmt.Sprintf("size.%v", Int)
  case float64:
    fn = fmt.Sprintf("size.%v", Float)
  case string:
    fn = fmt.Sprintf("size.%v", String)
  case bool:
    fn = fmt.Sprintf("size.%v", Bool)
  case *TCList:
    fn = fmt.Sprintf("size.%v", List)
  case *TCNumbers:
    fn = fmt.Sprintf("size.%v", Numbers)
  case *TCMatrix:
    fn = fmt.Sprintf("size.%v", Matrix)
  case *TCNeural:
    fn = fmt.Sprintf("size.%v", Neural)
  default:
    fn = fmt.Sprintf("size.%v", Any)
  }
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCSizeFun)
  }
  return nil
}


func init() {
  RegisterSizeCallback(String, TCStringSize)
  RegisterSizeCallback(Int, TCSizeSimple)
  RegisterSizeCallback(Float, TCSizeSimple)
  RegisterSizeCallback(Bool, TCSizeSimple)
  RegisterSizeCallback(Any, TCSizeSimple)
  RegisterSizeCallback(List, TCListSize)
  RegisterSizeCallback(Numbers, TCNumbersSize)
  RegisterSizeCallback(Matrix, TCMatrixSize)
  RegisterSizeCallback(Neural, TCNeuralSize)
}
