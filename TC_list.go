package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
)

type TCList struct {
  Q         deque.Deque
}

func (l *TCList) Add(x interface{}) {
  l.Q.PushBack(x)
}

func (l *TCList) Len() int {
  return l.Q.Len()
}

func (l *TCList) String() string {
  ret := "list[ "
  for x := 0; x < l.Q.Len(); x++ {
    e := l.Q.At(x)
    fun := GetConverterCallback(e)
    if fun == nil {
      continue
    }
    elem := fun(e, String)
    if elem == nil {
      continue
    }
    ret += fmt.Sprintf(" '%v' ", elem)
  }
  ret += " ]"
  return ret
}
