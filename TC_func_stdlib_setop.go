package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCSetOpFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperatorCommand("set", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  SetFunction("<-", TCSetOpFunction)
  SetFunction("set", TCSetOpFunction)
}
