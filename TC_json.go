package ThreadComputation
import (
  "time"
  "github.com/google/uuid"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
)

type TCJson struct {
  J        *gabs.Container
}

func tcFillJson(d *gabs.Container, q *deque.Deque) *gabs.Container {
  if (q.Len() % 2) != 0 {
    return nil
  }
  for q.Len() > 0 {
    k := q.PopFront()
    v := q.PopFront()
    if ! TCisSimple(v) {
      return nil
    }
    switch k.(type) {
    case string:
      d.Set(v, k.(string))
    default:
      return nil
    }
  }
  return d
}

func (tc *TCstate) Json() *TCJson {
  res := new(TCJson)
  res.J = gabs.New()
  return res
}

func (j *TCJson) String() string {
  return j.J.String()
}

func TCJsonFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.Json()
  res.J = tcFillJson(res.J, q)
  if res.J == nil {
    return nil, l.TC.MakeError("Error creating JSON")
  }
  return res, nil
}

func TCJsonValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.Json()
  res.J = tcFillJson(res.J, q)
  if res.J == nil {
    return nil, l.TC.MakeError("Error creating JSON")
  }
  return MakeValue(res, 100.0, 0), nil
}

func TCJsonUniqFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.Json()
  res.J = tcFillJson(res.J, q)
  if res.J == nil {
    return nil, l.TC.MakeError("Error creating JSON")
  }
  res.J.Set(time.Now().String(), "__time")
  res.J.Set(uuid.NewString(), "__id")
  return res, nil
}

func init() {
  SetCommand("json", TCJsonFunction)
  SetCommand("jsonUniq", TCJsonUniqFunction)
  SetCommand("JustJson", TCJsonValueFunction)
}
