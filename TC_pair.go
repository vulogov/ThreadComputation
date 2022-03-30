package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
)

type TCPair struct {
  X         interface{}
  Y         interface{}
}


func (tc *TCstate) Pair(x interface{}, y interface{}) *TCPair {
  res := new(TCPair)
  res.X = x
  res.Y = y
  return res
}

func (p *TCPair) String() string {
  return fmt.Sprintf("pair[ '%v' '%v']", p.X, p.Y)
}

func TCPairFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() != 2 {
    return nil, MakeError("pair[] requires 2 arguments")
  }
  x := q.PopFront()
  y := q.PopFront()
  res := l.TC.Pair(x,y)
  return res, nil
}

func TCPairValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() != 2 {
    return nil, MakeError("pair[] requires 2 arguments")
  }
  x := q.PopFront()
  y := q.PopFront()
  res := l.TC.Pair(x,y)
  return MakeValue(res, 100.0, 0), nil
}

func init() {
  SetCommand("pair", TCPairFunction)
  SetCommand("Pair", TCPairValueFunction)
}
