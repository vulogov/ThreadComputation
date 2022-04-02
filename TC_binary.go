package ThreadComputation
import (
  "github.com/gammazero/deque"
)

type TCBinary struct {
  Type        int
  D           []byte
}

func (b *TCBinary) String() string {
  return string(b.D)
}

func (b *TCBinary) Len() int {
  return len(b.D)
}

func (res *TCBinary) Add(d interface{}) *TCBinary {
  if d != nil {
    switch v := d.(type) {
    case string:
      res.Type = String
      res.D = []byte(v)
    case []byte:
      res.Type = String
      res.D = v
    case *TCBinary:
      res.Type = Binary
      res.D = append(res.D[:], v.D...)
    case *TCValue:
      switch v.Value.(type) {
      case string:
        res.Type = String
        res.D = []byte(v.Value.(string))
      case *TCBinary:
        res.Type = Binary
        res.D = append(res.D[:], v.Value.(*TCBinary).D...)
      }
    }
  }
  return res
}

func IsBinary(x interface{}) bool {
  switch x.(type) {
  case *TCBinary:
    return true
  }
  return false
}

func MakeBinary(d interface{}) *TCBinary {
  res := new(TCBinary)
  res.D = make([]byte, 0)
  res.Type = Unknown
  res.Add(d)
  return res
}

func TCBinaryFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      ReturnFromFunction(l, "bin", MakeBinary(e))
    }
  } else {
    return MakeBinary(nil), nil
  }
  return nil, nil
}

func TCBinaryValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      ReturnFromFunction(l, "bin", MakeValue(MakeBinary(e), 100.0, 0))
    }
  } else {
    return MakeValue(MakeBinary(nil), 100.0, 0), nil
  }
  return nil, nil
}

func init() {
  SetCommand("bin", TCBinaryFunction)
  SetCommand("JustBin", TCBinaryValueFunction)
}
