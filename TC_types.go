package ThreadComputation

import (
  "github.com/srfrog/dict"
  "gonum.org/v1/gonum/mat"
  "github.com/deckarep/golang-set"
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
  Error         = 97
  Simple        = 98
  Any           = 99
  Nil           = 100
)

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
  case *mat.Dense:
    return Matrix
  case *dict.Dict:
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
  case *TCError:
    return Error
  default:
    return Unknown
  }
}

func TCisSimple(x interface{}) bool {
  switch x.(type) {
  case int64,float64,string,bool:
    return true
  }
  return false
}
