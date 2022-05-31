package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCSampleFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("sample", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  SetFunction("sample", TCSampleFunction)
}
