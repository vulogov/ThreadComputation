package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCGetFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("get", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  SetFunction("->", TCGetFunction)
  SetFunction("get", TCGetFunction)
}
