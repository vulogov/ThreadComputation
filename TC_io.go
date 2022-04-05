package ThreadComputation
import (
  "fmt"
  "github.com/srfrog/dict"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
)

type TCIO struct {
  L          *TCExecListener
  Id          string
  D          *dict.Dict
}

func (r *TCIO) String() string {
  return fmt.Sprintf("io[ %v channel(s) ]", r.D.Len())
}

func IsIO(x interface{}) bool {
  switch x.(type) {
  case *TCIO:
    return true
  }
  return false
}

func (r *TCIO) Add(e interface{}) {
  switch v := e.(type) {
  case string:
    r.D.Set(uuid.NewString(), v)
  case *TCPair:
    switch key := v.X.(type) {
    case string:
      switch val := v.Y.(type) {
      case string:
        r.D.Set(key, val)
      }
    }
  }
}


func tcFillIO(io *TCIO, q *deque.Deque) *TCIO {
  for q.Len() > 0 {
    e := q.PopFront()
    io.Add(e)
  }
  return io
}

func MakeIO(l *TCExecListener, d interface{}) *TCIO {
  res := new(TCIO)
  res.L     = l
  res.Id    = uuid.NewString()
  res.D     = dict.New()
  if d != nil {
    switch d.(type) {
    case *deque.Deque:
      res = tcFillIO(res, d.(*deque.Deque))
    }
  }
  return res
}

func TCIOFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeIO(l, q), nil
}

func TCIOValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeValue(MakeIO(l, q), 100.0, 0), nil
}

func init() {
  SetCommand("io", TCIOFunction)
  SetCommand("JustIO", TCIOValueFunction)
}
