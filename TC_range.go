package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
)

type TCRange struct {
  X          float64
  Y          float64
  Step       float64
}

func toFloatIfPossible(v interface{}) interface{} {
  switch v.(type) {
  case float64:
    return v.(float64)
  case int64:
    return float64(v.(int64))
  }
  return nil
}

func (r *TCRange) String() string {
  return fmt.Sprintf("range[%v...%v (%v)]", r.X, r.Y, r.Step)
}

func TCRangeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var x, y, step float64
  x = 0.0
  y = 1.0
  step = 1.0
  if q.Len() == 1 {
    e := q.PopFront()
    res := toFloatIfPossible(e)
    if res != nil {
      y = res.(float64)
    }
  } else if q.Len() == 2 {
    e1 := q.PopFront()
    e2 := q.PopFront()
    res := toFloatIfPossible(e1)
    if res != nil {
      x = res.(float64)
    }
    res = toFloatIfPossible(e2)
    if res != nil {
      y = res.(float64)
    }
  } else if q.Len() > 2 {
    e1 := q.PopFront()
    e2 := q.PopFront()
    e3 := q.PopFront()
    res := toFloatIfPossible(e1)
    if res != nil {
      x = res.(float64)
    }
    res = toFloatIfPossible(e2)
    if res != nil {
      y = res.(float64)
    }
    res = toFloatIfPossible(e3)
    if res != nil {
      step = res.(float64)
    }
  }
  if x > y {
    return nil, l.TC.MakeError("Ranges in range is invalid: %v %v", x, y)
  }
  res := new(TCRange)
  res.X = x
  res.Y = y
  res.Step = step
  return res, nil
}

func init() {
  SetCommand("range", TCRangeFunction)
}
