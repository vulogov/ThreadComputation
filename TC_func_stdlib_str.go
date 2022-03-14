package ThreadComputation

import (
  "fmt"
  "strings"
  "github.com/gammazero/deque"
)





func TCStrFormatFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var args []interface{}

  if q.Len() > 0 {
    f := q.PopFront()
    switch f.(type) {
    case string:
      for q.Len() > 0 {
        v := q.PopFront()
        fun := GetConverterCallback(v)
        if fun == nil {
          continue
        }
        args = append(args, fun(v, String))
      }
      ReturnFromFunction(l, "str.format", fmt.Sprintf(f.(string), args...))
    }
  }
  return nil, nil
}

func TCStrStringFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
      s := q.PopFront()
      switch s.(type) {
      case string:
        switch name {
        case "str.strip":
          ReturnFromFunction(l, "str.strip", strings.Trim(s.(string), "\n \t\r"))
        case "str.title":
          ReturnFromFunction(l, "str.title", strings.Title(s.(string)))
        case "str.lower":
          ReturnFromFunction(l, "str.lower", strings.ToLower(s.(string)))
        case "str.upper":
          ReturnFromFunction(l, "str.upper", strings.ToUpper(s.(string)))
        default:
          break
        }
      default:
        break
      }
    }
    return nil, nil
}


func init() {
  SetFunction("str.format", TCStrFormatFunction)
  SetFunction("str.strip", TCStrStringFunction)
  SetFunction("str.title", TCStrStringFunction)
  SetFunction("str.lower", TCStrStringFunction)
  SetFunction("str.upper", TCStrStringFunction)
}
