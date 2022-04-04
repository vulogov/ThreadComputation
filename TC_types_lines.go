package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCLinesTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    res := MakeLines(q)
    return res, nil
  }
  return MakeType(Lines), nil
}

func init() {
  SetCommand("Lines", TCLinesTypeFunction)
}
