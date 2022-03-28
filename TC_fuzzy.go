package ThreadComputation

import (
  "fmt"
  "time"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

type TCValue struct {
  Value      interface{}
  Type       int
  Stamp      time.Time
  TTL        int
  P          float32
}

func MakeValue(v interface{}, p float32, ttl int) *TCValue {
  res := new(TCValue)
  res.Stamp = time.Now()
  res.Value = v
  res.P     = p
  res.TTL   = ttl
  log.Debugf("Creating fuzzy value of type %T %v", res.Value, res.P)
  return res
}

func (tc *TCstate) Value(v interface{}, p float32, ttl int) *TCValue {
  return MakeValue(v, p, ttl)
}

func (v *TCValue) String() string {
  ofun := GetConverterCallback(v.Value)
  if ofun == nil {
    return fmt.Sprintf("Value[%v %v]", v.Value, v.P)
  }
  out := ofun(v.Value, String)
  if out == nil {
    out = fmt.Sprintf("%v", v.Value)
  }
  return fmt.Sprintf("Value[value=%v probability=%v%%]", out, v.P)
}

func (v *TCValue) IsTrue() bool {
  switch v.Value.(type) {
  case bool:
    return v.Value.(bool)
  }
  return false
}

func tcvalueGetValue(l *TCExecListener, q *deque.Deque) interface{} {
  if q.Len() == 0 {
    if l.TC.Ready() {
      v := l.TC.Get()
      return v
    }
  }
  if q.Len() >= 1 {
    v := q.PopFront()
    return v
  }
  return nil
}

func tcvalueGetP(l *TCExecListener, q *deque.Deque) float32 {
  if q.Len() == 0 {
    return float32(100.0)
  }
  if q.Len() > 0 {
    v := q.PopFront()
    switch v.(type) {
    case float64:
      return float32(v.(float64))
    }
  }
  return float32(100.0)
}

func TCValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  v := tcvalueGetValue(l, q)
  if v == nil {
    return nil, l.TC.MakeError("Can not get value from context")
  }
  p := tcvalueGetP(l, q)
  return l.TC.Value(v, p, 0), nil
}

func TCVFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 1 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      return e.(*TCValue).Value, nil
    default:
      return nil, l.TC.MakeError("V[] expecting to have a value")
    }
  }
  if q.Len() > 1 {
    e := q.PopFront()
    data := q.PopFront()
    switch e.(type) {
    case *TCValue:
      e.(*TCValue).Value = data
      return e, nil
    }
  }
  return nil, l.TC.MakeError("Incorrect context for V[]")
}

func TCJustValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  v := tcvalueGetValue(l, q)
  if v == nil {
    return nil, l.TC.MakeError("Can not get value from context")
  }
  return l.TC.Value(v, float32(100.0), 0), nil
}

func TCNeverValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  v := tcvalueGetValue(l, q)
  if v == nil {
    return nil, l.TC.MakeError("Can not get value from context")
  }
  return l.TC.Value(v, float32(0.0), 0), nil
}

func TCLikelyValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  v := tcvalueGetValue(l, q)
  if v == nil {
    return nil, l.TC.MakeError("Can not get value from context")
  }
  return l.TC.Value(v, float32(50.0), 0), nil
}

func TCValueMakeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var p float32
  e := q.PopFront()
  if q.Len() == 0 {
    p = 100.0
  } else {
    pp := q.PopFront()
    switch pp.(type) {
    case float64:
      p = float32(pp.(float64))
    }
  }
  return l.TC.Value(e, p, 0), nil
}

func TCSureValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return l.TC.Value(true, 100.0, 0), nil
}

func TCNotSureValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return l.TC.Value(false, 100.0, 0), nil
}


func init() {
  SetCommand("Value", TCValueFunction)
  SetCommand("Just", TCJustValueFunction)
  SetCommand("Never", TCNeverValueFunction)
  SetCommand("Likely", TCLikelyValueFunction)
  SetCommand("Sure", TCSureValueFunction)
  SetCommand("NotSure", TCNotSureValueFunction)
  SetFunction("V", TCVFunction)
  SetFunction("Make", TCValueMakeFunction)
}
