package ThreadComputation

import (
  "fmt"
  "strings"
  "strconv"
  "github.com/deckarep/golang-set"
  conv "github.com/cstockton/go-conv"
  log "github.com/sirupsen/logrus"
)



type TCConvertFun func(interface{}, int) interface{}

func TCStringConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case string:
    switch to_type {
    case String:
      if strings.HasPrefix(e, "'") &&  strings.HasSuffix(e, "'")  {
        e = strings.TrimPrefix(e, "'")
        e = strings.TrimSuffix(e, "'")
        return e
      } else if strings.HasPrefix(e, "\"") &&  strings.HasSuffix(e, "\"") {
        e = strings.TrimPrefix(e, "\"")
        e = strings.TrimSuffix(e, "\"")
        return e
      }
    case Int:
      ret, err := strconv.ParseInt(e, 10, 64)
      if err == nil {
        return ret
      }
    case Float:
      ret, err := strconv.ParseFloat(e, 64)
      if err == nil {
        return ret
      }
    case Bool:
      ret, err := conv.Bool(data)
      if err == nil {
        return ret
      }
    }
  }
  return nil
}

func TCIntConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case int64:
    switch to_type {
    case String:
      return fmt.Sprintf("%v", e)
    case Int:
      ret, err := conv.Int64(data)
      if err == nil {
        return ret
      }
    case Float:
      ret, err := conv.Float64(float64(e))
      if err == nil {
        return ret
      }
    case Bool:
      if e == 1 {
        return true
      } else {
        return false
      }
    }
  }
  return nil
}

func TCFloatConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case float64:
    switch to_type {
    case String:
      return fmt.Sprintf("%v", e)
    case Int:
      ret, err := conv.Int64(int64(data.(float64)))
      if err == nil {
        return ret
      }
    case Float:
      ret, err := conv.Float64(float64(e))
      if err == nil {
        return ret
      }
    case Bool:
      if e == 1 {
        return true
      } else {
        return false
      }
    }
  }
  return nil
}

func TCBoolConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case bool:
    switch to_type {
    case String:
      return fmt.Sprintf("%v", e)
    case Int:
      if e {
        return int64(1)
      } else {
        return int64(0)
      }
    case Float:
      if e {
        return float64(1)
      } else {
        return float64(0)
      }
    case Bool:
      return e
    }
  }
  return nil
}

func TCListConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCList:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCSetConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case mapset.Set:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCRangeConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCRange:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCNoneConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCNone:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCAnythingConvert(data interface{}, to_type int) interface{} {
  return fmt.Sprintf("%v", data)
}

func TCErrorConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCError:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCNumbersConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCNumbers:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCMatrixConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCMatrix:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCDictConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCDict:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCPairConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCPair:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCJsonConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCJson:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCNeuralConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCNeural:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func RegisterConvertCallback(from_type int, fun TCConvertFun) {
  fname := fmt.Sprintf("convert.%v", from_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func GetConverterCallback(x interface{}) TCConvertFun {
  var fn string
  log.Debugf("Request converter for: %T", x)
  switch x.(type) {
  case int64:
    fn = fmt.Sprintf("convert.%v", Int)
  case float64:
    fn = fmt.Sprintf("convert.%v", Float)
  case string:
    fn = fmt.Sprintf("convert.%v", String)
  case bool:
    fn = fmt.Sprintf("convert.%v", Bool)
  case *TCList:
    fn = fmt.Sprintf("convert.%v", List)
  case mapset.Set:
    fn = fmt.Sprintf("convert.%v", Set)
  case *TCCode:
    fn = fmt.Sprintf("convert.%v", Code)
  case *TCFunRef:
    fn = fmt.Sprintf("convert.%v", Ref)
  case *TCRange:
    fn = fmt.Sprintf("convert.%v", Range)
  case *TCNone:
    fn = fmt.Sprintf("convert.%v", None)
  case *TCNumbers:
    fn = fmt.Sprintf("convert.%v", Numbers)
  case *TCMatrix:
    fn = fmt.Sprintf("convert.%v", Matrix)
  case *TCDict:
    fn = fmt.Sprintf("convert.%v", Dict)
  case *TCPair:
    fn = fmt.Sprintf("convert.%v", Pair)
  case *TCJson:
    fn = fmt.Sprintf("convert.%v", Json)
  case *TCNeural:
    fn = fmt.Sprintf("convert.%v", Neural)
  case *TCError:
    fn = fmt.Sprintf("convert.%v", Error)
  default:
    fn = fmt.Sprintf("convert.%v", Any)
  }
  fun, ok := Callbacks.Load(fn)
  if ok {
    log.Debugf("Returning converter: %v", fn)
    return fun.(TCConvertFun)
  }
  return nil
}

func GetStringConverter() TCConvertFun {
  fun, ok := Callbacks.Load(fmt.Sprintf("convert.%v", String))
  if ok {
    return fun.(TCConvertFun)
  }
  return nil
}

func init() {
  RegisterConvertCallback(String, TCStringConvert)
  RegisterConvertCallback(Int, TCIntConvert)
  RegisterConvertCallback(Float, TCFloatConvert)
  RegisterConvertCallback(Bool, TCBoolConvert)
  RegisterConvertCallback(Any, TCAnythingConvert)
  RegisterConvertCallback(List, TCListConvert)
  RegisterConvertCallback(Set, TCSetConvert)
  RegisterConvertCallback(Range, TCRangeConvert)
  RegisterConvertCallback(Error, TCErrorConvert)
  RegisterConvertCallback(None, TCNoneConvert)
  RegisterConvertCallback(Numbers, TCNumbersConvert)
  RegisterConvertCallback(Matrix, TCMatrixConvert)
  RegisterConvertCallback(Dict, TCDictConvert)
  RegisterConvertCallback(Pair, TCPairConvert)
  RegisterConvertCallback(Json, TCJsonConvert)
  RegisterConvertCallback(Neural, TCNeuralConvert)
}
