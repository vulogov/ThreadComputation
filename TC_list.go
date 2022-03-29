package ThreadComputation
import (
  "fmt"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
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
    switch e.(type) {
    case string:
      ret += fmt.Sprintf(" '%v' ", e)
      continue
    }
    fun := GetConverterCallback(e)
    if fun == nil {
      log.Debugf("list convert: Can not find convert function for %T", e)
      continue
    }
    elem := fun(e, String)
    if elem == nil {
      log.Debugf("list convert: convert function return nil for %T", e)
      continue
    }
    ret += fmt.Sprintf(" '%v' ", elem)
  }
  ret += " ]"
  return ret
}
