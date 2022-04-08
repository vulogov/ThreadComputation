package ThreadComputation

import (
  "github.com/gammazero/deque"
)


func TCExportFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {

  return nil, nil
}

func init() {
  SetCommand("export", TCExportFunction)
}
