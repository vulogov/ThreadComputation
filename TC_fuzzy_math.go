package ThreadComputation

import (
  "strings"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

func tcFuzzyMathFindMatch(q *deque.Deque, igx interface{}) (int, float32, interface{}) {
  var p float32
  var v interface{}
  i := 0
  p = 0.0
  for x := 0; x < q.Len(); x++ {
    if igx != nil {
      if igx.(int) == x {
        continue
      }
    }
    e := q.At(x)
    switch e.(type) {
    case *TCValue:
      if e.(*TCValue).P > p {
        p = e.(*TCValue).P
        i = x
        v = e
      }
    default:
      if p < 100.0 {
        p = 100.0
        i = x
        v = e
      }
    }
    if p == 100.0 {
      return i, p, v
    }
  }
  if v != nil {
    return i, p, v
  }
  return 0, 0.0, nil
}

func TCFuzzyMathFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() < 2 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      return e, nil
    }
    return nil, l.TC.MakeError("Fuzzy math operator expecting Value[] in stack")
  }
  operator := strings.TrimPrefix(name, "^")
  i, p1, x := tcFuzzyMathFindMatch(q, nil)
  if x == nil {
    return nil, l.TC.MakeError("Fuzzy math operator did not find value for x")
  }
  _, p2, y := tcFuzzyMathFindMatch(q, i)
  if y == nil {
    return nil, l.TC.MakeError("Fuzzy math operator did not find value for y")
  }
  fun := GetMathCallback(x, y)
  if fun == nil {
    return nil, l.TC.MakeError("Fuzzy math operator did not find math callback for %T %T", x, y)
  }
  res := fun(operator, x, y)
  if res == nil {
    return nil, l.TC.MakeError("Fuzzy math operator for %T %T return nil", x, y)
  }
  p := (p1+p2)/2
  log.Debugf("Fuzzy operator: %v %v %v = %v probability=%v", x, operator, y, res, p)
  switch res.(type) {
  case *TCValue:
    return res, nil
  default:
    return MakeValue(res, p, 0), nil
  }
  return nil, nil
}

func init() {
  SetFunction("^+", TCFuzzyMathFunction)
  SetFunction("^*", TCFuzzyMathFunction)
  SetFunction("^-", TCFuzzyMathFunction)
  SetFunction("^/", TCFuzzyMathFunction)
}
