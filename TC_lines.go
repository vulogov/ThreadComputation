package ThreadComputation
import (
  "fmt"
  "regexp"
  "github.com/gammazero/deque"
)

var TCEndOfLine = regexp.MustCompile(`\n+`)

type TCLines struct {
  Q         deque.Deque
}

func MakeLinesFromValue(d interface{}) *TCLines {
  switch d.(type) {
  case *TCValue:
    return MakeLinesFromValue(d.(*TCValue).Value)
  }
  res := new(TCLines)
  switch v := d.(type) {
  case string:
    for _, elem := range TCEndOfLine.FindAllString(v, -1) {
      res.Add(elem)
    }
  case *TCBinary:
    return MakeLinesFromValue(string(v.Raw()))
  }
  return res
}

func MakeLines(q *deque.Deque) *TCLines {
  res := new(TCLines)
  for q.Len() > 0 {
    e := q.PopFront()
    res.Add(e)
  }
  return res
}

func IsLines(x interface{}) bool {
  switch x.(type) {
  case *TCLines:
    return true
  }
  return false
}

func (l *TCLines) Add(x interface{}) {
  val := fmt.Sprintf("%v", x)
  switch x.(type) {
  case string:
    val = x.(string)
  case *TCValue:
    l.Add(x.(*TCValue).Value)
  case *TCBinary:
    val = string(x.(*TCBinary).D)
  default:
    fun := GetConverterCallback(x)
    if fun != nil {
      res := fun(x, String)
      if res != nil {
        val = res.(string)
      }
    }
  }
  l.Q.PushBack(val)
}

func (l *TCLines) Len() int {
  return l.Q.Len()
}

func (l *TCLines) String() string {
  out := ""
  for i := 0; i < l.Len(); i++ {
    out += fmt.Sprintf("%v\n", l.Q.At(i))
  }
  return out
}

func TCLinesFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeLines(q)
  return res, nil
}

func TCJustLinesFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return MakeValue(MakeLines(q), 100.0, 0), nil
}

func init() {
  SetCommand("lines", TCLinesFunction)
  SetCommand("JustLines", TCJustLinesFunction)
}
