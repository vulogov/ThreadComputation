package ThreadComputation

import (
  "github.com/gammazero/deque"
)


func NowFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return nil, nil
}



func initStdlibTime() {
  SetFunction("now", NowFunction)
}
