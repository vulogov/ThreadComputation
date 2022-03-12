package ThreadComputation

import (
  "fmt"
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




func init() {
  SetFunction("str.format", TCStrFormatFunction)
}
