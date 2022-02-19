package ThreadComputation

import (
  "fmt"
  "errors"
  "strings"
  "github.com/gammazero/deque"
  "github.com/agnivade/levenshtein"
)

func TCDistanceFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var e1 interface{}
  var e2 interface{}

  if l.TC.Ready() && q.Len() == 0 {
    if l.TC.Res.Len() >= 2 {
      e1 = l.TC.Get()
      e2 = l.TC.Get()
    } else {
      return nil, errors.New("expecting to have two elements for comparing")
    }
  } else if l.TC.Ready() && q.Len() == 1  {
    e1 = l.TC.Get()
    e2 = q.PopFront()
  } else if q.Len() == 2 {
    e1 = q.PopFront()
    e2 = q.PopFront()
  } else {
    return nil, errors.New("Inconclusive context for comparing")
  }

  switch e1.(type) {
  case string:
    switch e2.(type) {
    case string:
      d := levenshtein.ComputeDistance(e1.(string), e2.(string))
      return d, nil
    }
  }
  return nil, errors.New(fmt.Sprintf("Unknown context for distance operator: %v(%T) %v(%T)", e1, e1, e2, e2))
}

func StringFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      s := q.PopFront()
      switch s.(type) {
      case string:
        switch l.TC.CurrentFunctionName() {
        case "strip":
          ReturnFromFunction(l, "strip", strings.Trim(s.(string), "\n \t\r"))
        case "title":
          ReturnFromFunction(l, "title", strings.Title(s.(string)))
        case "lower":
          ReturnFromFunction(l, "lower", strings.ToLower(s.(string)))
        case "upper":
          ReturnFromFunction(l, "upper", strings.ToUpper(s.(string)))
        default:
          break
        }
      default:
        break
      }
    }
    return nil, nil
  } else {
    if l.TC.Ready() {
      s := l.TC.Get()
      switch s.(type) {
      case string:
        switch l.TC.CurrentFunctionName() {
        case "strip":
          return strings.Trim(s.(string), "\n \t\r"), nil
        case "title":
          return strings.Title(s.(string)), nil
        case "upper":
          return strings.ToUpper(s.(string)), nil
        case "lower":
          return strings.ToLower(s.(string)), nil
        }
      }
    }
  }
  return nil, errors.New("strip function did not discover proper context")
}



func initStdlibString() {
  SetFunction("strip", StringFunction)
  SetFunction("title", StringFunction)
  SetFunction("lower", StringFunction)
  SetFunction("upper", StringFunction)
  SetFunction("distance", TCDistanceFunction)
}
