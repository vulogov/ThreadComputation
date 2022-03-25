package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func blockConditionalFunction(l *TCExecListener, name string, code string, cond bool) interface{} {
  if l.TC.Ready() {
    e := l.TC.Get()
    switch e.(type) {
    case bool:
      if e.(bool) == cond {
        log.Debugf("stdlib conditional %v matches %v condition", name, cond)
        l.TC.Eval(code)
      }
    default:
      l.TC.Set(e)
    }
  }
  return nil
}

func BlockConditionalFalse(l *TCExecListener, name string, code string) interface{} {
  return blockConditionalFunction(l, name, code, false)
}

func BlockConditionalTrue(l *TCExecListener, name string, code string) interface{} {
  return blockConditionalFunction(l, name, code, true)
}

func init() {
  RegisterBlockCallback("iftrue", BlockConditionalTrue)
  RegisterBlockCallback("iffalse", BlockConditionalFalse)
}
