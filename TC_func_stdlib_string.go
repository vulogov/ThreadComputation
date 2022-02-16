package ThreadComputation

import (
  "errors"
  "strings"
  "github.com/gammazero/deque"
)


func StringFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      s := q.PopFront()
      switch s.(type) {
      case string:
        switch l.TC.CurrentFunctionName() {
        case "strip":
          l.TC.Res.Set(strings.Trim(s.(string), "\n \t\r"))
        case "title":
          l.TC.Res.Set(strings.Title(s.(string)))
        case "lower":
          l.TC.Res.Set(strings.ToLower(s.(string)))
        case "upper":
          l.TC.Res.Set(strings.ToUpper(s.(string)))
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
}
