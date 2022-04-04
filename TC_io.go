package ThreadComputation
import (
  "fmt"
  "github.com/srfrog/dict"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
)

type TCIO struct {
  Id          string
  D          *dict.Dict
}

func (r *TCIO) String() string {
  return fmt.Sprintf("io[set at %v]", r.D.Len())
}

func IsIO(x interface{}) bool {
  switch x.(type) {
  case *TCIO:
    return true
  }
  return false
}

func MakeIO() *TCIO {
  res := new(TCIO)
  res.Id    = uuid.NewString()
  res.D     = dict.New()
  return res
}

func TCIOFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeIO(), nil
}

func TCIOValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeValue(MakeIO(), 100.0, 0), nil
}

func init() {
  SetCommand("io", TCIOFunction)
  SetCommand("JustIO", TCIOValueFunction)
}
