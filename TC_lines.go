package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
)

type TCLines struct {
  Q         deque.Deque
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
  return fmt.Sprintf("lines[ %v line(s) ]", l.Len())
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
