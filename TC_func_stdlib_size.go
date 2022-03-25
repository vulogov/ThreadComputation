package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCSizeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopBack()
    fun := GetSizeCallback(e)
    if fun != nil {
      ReturnFromFunction(l, "size", fun(e))
    } else {
      return nil, l.TC.MakeError("Can not determine size of %T", e)
    }
  }
  return nil, nil
}

func init() {
  SetFunction("size", TCSizeFunction)
}
