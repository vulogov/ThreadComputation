package ThreadComputation

import (
  "github.com/gammazero/deque"
)


func TCCarFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("car", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func TCCdrFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("cdr", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  SetCommand("car", TCCarFunction)
  SetCommand("cdr", TCCdrFunction)
}
