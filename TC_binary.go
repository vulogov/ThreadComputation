package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

type TCBinary struct {
  Type        int
  D           []byte
}

func (b *TCBinary) String() string {
  out := fmt.Sprintf("binary(%v)[ ", TypeToStr(b.Type))
  for i := 0; i<len(b.D); i++ {
    out += fmt.Sprintf("%U(%q) ", b.D[i], b.D[i])
  }
  out += " ]"
  return out
}

func (b *TCBinary) Raw() string {
  return string(b.D)
}

func (b *TCBinary) Slice(s int, l int) *TCBinary {
  if s+l <= len(b.D) {
    res := MakeBinary(b.D[s:s+l])
    return res
  }
  return nil
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
    default:
      log.Debugf("binary-add: No suitable converter")
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
  cfun := GetConverterCallback(d)
  if cfun != nil {
    res := cfun(d, Binary)
    if res != nil {
      switch res.(type) {
      case *TCBinary:
        log.Debugf("make-binary: done using standard converter for %T", d)
        return res.(*TCBinary)
      }
    } else {
      log.Debugf("make-binary: standard converter for %T returned nil", d)
    }
  } else {
    log.Debugf("make-binary: no standard converter for %T", d)
  }
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
      log.Debugf("Passed on bin[] input: %v", e)
      ReturnFromFunction(l, "bin", MakeBinary(e))
    }
    return nil, nil
  }
  return MakeBinary(nil), nil
}

func TCBinaryValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      ReturnFromFunction(l, "bin", MakeValue(MakeBinary(e), 100.0, 0))
    }
    return nil, nil
  }
  return MakeValue(MakeBinary(nil), 100.0, 0), nil
}

func init() {
  SetCommand("bin", TCBinaryFunction)
  SetCommand("JustBin", TCBinaryValueFunction)
}
