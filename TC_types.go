package ThreadComputation

import (
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)

const (
  Unknown   int = 0
  Int           = 1
  Float         = 2
  String        = 3
  Bool          = 4
  List          = 5
  Matrix        = 6
  Dict          = 7
  Set           = 8
  Ref           = 9
  Code          = 10
  Value         = 11
  Range         = 12
  None          = 13
  Numbers       = 14
  Pair          = 15
  Json          = 16
  Neural        = 17
  Data          = 18
  Time          = 19
  Binary        = 20
  Context       = 21
  Lines         = 22
  IO            = 23
  Exported      = 24
  Iterator      = 95
  SType         = 96
  Error         = 97
  Simple        = 98
  Any           = 99
  Nil           = 100
)

type Type struct {
  T      int
}

func (t *Type) String() string {
  return TypeToStr(t.T)
}

func IsType(x interface{}) bool {
  switch x.(type) {
  case *Type:
    return true
  }
  return false
}

func MakeType(t interface{}) *Type {
  res := new(Type)
  res.T    = Unknown
  switch t.(type) {
  case int:
    res.T = t.(int)
  default:
    res.T = TCType(t)
  }
  return res
}

func SetExternalTypeHandlers(fun1 TCExtType, fun2 TCExtTypeStr) {
  extType = fun1
  extTypeStr = fun2
}

func TypeToStr(t interface{}) string {
  switch t.(type) {
  case int:
    switch t {
    case Int:
      return "Int"
    case Float:
      return "Float"
    case String:
      return "String"
    case Bool:
      return "Bool"
    case List:
      return "List"
    case Matrix:
      return "Matrix"
    case Dict:
      return "Dict"
    case Set:
      return "Set"
    case Ref:
      return "Ref"
    case Code:
      return "Code"
    case Value:
      return "Value"
    case Range:
      return "Range"
    case None:
      return "None"
    case Numbers:
      return "Numbers"
    case Pair:
      return "Pair"
    case Json:
      return "Json"
    case Neural:
      return "Neural"
    case Data:
      return "Data"
    case Binary:
      return "Binary"
    case Context:
      return "Context"
    case Iterator:
      return "Iterator"
    case Lines:
      return "Lines"
    case IO:
      return "IO"
    case Exported:
      return "Exported"
    case Error:
      return "Error"
    case SType:
      return "Type"
    default:
      if extTypeStr != nil {
        return extTypeStr(t)
      }
    }
  }
  return TypeToStr(TCType(t))
}

func TCType(x interface{}) int {
  if x == nil {
    return Nil
  }
  switch x.(type) {
  case int64:
    return Int
  case float64:
    return Float
  case string:
    return String
  case bool:
    return Bool
  case *TCList:
    return List
  case *TCDict:
    return Dict
  case mapset.Set:
    return Set
  case *TCFunRef:
    return Ref
  case *TCCode:
    return Code
  case *TCValue:
    return Value
  case *TCRange:
    return Range
  case *TCNone:
    return None
  case *TCNumbers:
    return Numbers
  case *TCMatrix:
    return Matrix
  case *TCPair:
    return Pair
  case *TCJson:
    return Json
  case *TCNeural:
    return Neural
  case *TCData:
    return Data
  case *TCIterator:
    return Iterator
  case *TCBinary:
    return Binary
  case *TCCtx:
    return Context
  case *TCLines:
    return Lines
  case *TCIO:
    return IO
  case *TCExported:
    return Exported
  case *TCError:
    return Error
  case *Type:
    return x.(*Type).T
  default:
    if extType != nil {
      return extType(x)
    }
  }
  return Unknown
}

func TCisSimple(x interface{}) bool {
  switch x.(type) {
  case int64,float64,string,bool:
    return true
  }
  return false
}

func TCBoolTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case bool:
        ReturnFromFunction(l, "Int", e)
      }
    }
    return nil, nil
  }
  return MakeType(Bool), nil
}

func TCIntFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case int64:
        ReturnFromFunction(l, "Int", e)
      case float64:
        ReturnFromFunction(l, "Int", int64(e.(float64)))
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case int64:
          ReturnFromFunction(l, "Int", e)
        }
      default:
        fun := GetConverterCallback(e)
        if fun == nil {
          return nil, l.TC.MakeError("Error getting converter for Int")
        }
        res := fun(e, Int)
        if res == nil {
          return nil, l.TC.MakeError("Error converting for Int")
        }
        ReturnFromFunction(l, "Int", res)
      }
    }
    return nil, nil
  }
  return MakeType(Int), nil
}

func TCFloatFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case float64:
        ReturnFromFunction(l, "Float", e)
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case float64:
          ReturnFromFunction(l, "Float", e)
        }
      case int64:
        ReturnFromFunction(l, "Float", float64(e.(int64)))
      default:
        fun := GetConverterCallback(e)
        if fun == nil {
          return nil, l.TC.MakeError("Error getting converter for Float")
        }
        res := fun(e, Float)
        if res == nil {
          return nil, l.TC.MakeError("Error converting for Float")
        }
        ReturnFromFunction(l, "Float", res)
      }
    }
    return nil, nil
  }
  return MakeType(Float), nil
}

func TCStringFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case string:
        ReturnFromFunction(l, "String", e)
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case string:
          ReturnFromFunction(l, "String", e)
        }
      default:
        fun := GetConverterCallback(e)
        if fun != nil {
          res := fun(e, String)
          if res == nil {
            ReturnFromFunction(l, "String", res)
          }
        }
      }
    }
    return nil, nil
  }
  return MakeType(String), nil
}

func TCTimeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeType(Time), nil
}

func TCListTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    res := new(TCList)
    for q.Len() > 0 {
      e := q.PopFront()
      res.Q.PushBack(e)
    }
    return res, nil
  }
  return MakeType(List), nil
}

func TCNumbersTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    res := MakeNumbers()
    for q.Len() > 0 {
      e := q.PopFront()
      res.Add(e)
    }
    return res, nil
  }
  return MakeType(Numbers), nil
}

func TCToTypeConverter(v interface{}) interface{} {
  return MakeType(v)
}

func TCToTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("totype", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  SetCommand("Int", TCIntFunction)
  SetCommand("Float", TCFloatFunction)
  SetCommand("String", TCStringFunction)
  SetCommand("Time", TCTimeFunction)
  SetCommand("List", TCListTypeFunction)
  SetCommand("Numbers", TCNumbersTypeFunction)
  SetCommand("Bool", TCBoolTypeFunction)
  SetFunction("type", TCToTypeFunction)
  RegisterFunctionCallback("totype", Any, TCToTypeConverter)
}
