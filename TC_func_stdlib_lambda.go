package ThreadComputation

import (
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)


func BlockLambda(l *TCExecListener, name string, code string) interface{} {
  log.Debugf("Creating pointer on lambda function: %v", name)
  ReturnFromFunction(l, "λ", l.TC.NewRef(name, new(deque.Deque)))
  return nil
}

func init() {
  RegisterBlockCallback("lambda", BlockLambda)
  RegisterBlockCallback("λ", BlockLambda)
}
