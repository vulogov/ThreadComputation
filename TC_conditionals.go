package ThreadComputation
import (
  log "github.com/sirupsen/logrus"
)

func (tc *TCstate) IsInConditional() bool {
  if tc.CStack.Len() > 0 {
    return true
  }
  return false
}

func (tc *TCstate) IfConditional(name string) bool {
  if tc.CStack.Len() > 0 {
    if tc.CStack.Front().(string) == name {
      return true
    }
  }
  return false
}

func (tc *TCstate) EndConditional(name string) bool {
  if tc.CStack.Len() > 0 {
    if tc.CStack.Front().(string) == name {
      log.Debugf("END condition: %v", name)
      tc.CStack.PopFront()
      return true
    }
  }
  log.Debugf("CONTINUE condition: %v", name)
  return false
}

func (tc *TCstate) BeginConditional(name string) {
  log.Debugf("BEGIN condition: %v", name)
  tc.CStack.PushFront(name)
}
