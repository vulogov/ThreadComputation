package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func InCode(l *TCExecListener, name string, code string) interface{} {
  log.Debugf("in: %v", name)
  l.TC.AddContext(nil)
  l.TC.SetContext("inContext", true)
  l.TC.Eval(code)
  l.TC.DelContext()
  return nil
}


func init() {
  RegisterBlockCallback("in", InCode)
}
