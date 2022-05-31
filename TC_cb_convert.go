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

func TCDataConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCData:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCTypeConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *Type:
    switch to_type {
    case String:
      return e.String()
    case Int:
      return int64(e.T)
    }
  }
  return nil
}

func TCIteratorConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCIterator:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCContextConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCCtx:
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
  fn = fmt.Sprintf("convert.%v", TCType(x))
  fun, ok := Callbacks.Load(fn)
  if ok {
    log.Debugf("Returning converter: %v", fn)
    return fun.(TCConvertFun)
  }
  fn = fmt.Sprintf("convert.%v", Any)
  fun, ok = Callbacks.Load(fn)
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
  RegisterConvertCallback(Any, TCAnythingConvert)
  RegisterConvertCallback(Set, TCSetConvert)
  RegisterConvertCallback(Range, TCRangeConvert)
  RegisterConvertCallback(Error, TCErrorConvert)
  RegisterConvertCallback(None, TCNoneConvert)
  RegisterConvertCallback(Matrix, TCMatrixConvert)
  RegisterConvertCallback(Dict, TCDictConvert)
  RegisterConvertCallback(Pair, TCPairConvert)
  RegisterConvertCallback(Json, TCJsonConvert)
  RegisterConvertCallback(Neural, TCNeuralConvert)
  RegisterConvertCallback(Data, TCDataConvert)
  RegisterConvertCallback(SType, TCTypeConvert)
  RegisterConvertCallback(Iterator, TCIteratorConvert)
  RegisterConvertCallback(Context, TCContextConvert)
}
