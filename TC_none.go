package ThreadComputation
import (
  "fmt"
  "time"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
)

type TCNone struct {
  Id          string
  Stamp       time.Time
}

func (r *TCNone) String() string {
  return fmt.Sprintf("none[set at %v]", r.Stamp.Format(time.RFC3339))
}

func IsNone(x interface{}) bool {
  switch x.(type) {
  case *TCNone:
    return true
  }
  return false
}

func MakeNone() *TCNone {
  res := new(TCNone)
  res.Id    = uuid.NewString()
  res.Stamp = time.Now()
  return res
}

func TCNoneFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeNone(), nil
}

func TCNoneValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := new(TCNone)
  res.Id    = uuid.NewString()
  res.Stamp = time.Now()
  return MakeValue(res, 100.0, 0), nil
}

func init() {
  SetCommand("none", TCNoneFunction)
  SetCommand("JustNone", TCNoneValueFunction)
}
