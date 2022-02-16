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
        }
      }
    }
  }
  return nil, errors.New("strip function did not discover proper context")
}



func initStdlibString() {
  SetFunction("strip", StringFunction)
}
