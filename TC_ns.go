package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)

func (l *TCExecListener) EnterNs(c *parser.NsContext) {
  if l.TC.Errors() > 0 {
    return
  }
  ns_name := c.GetNsname().GetText()
  log.Debugf("Entering NS: %v", ns_name)
  if l.TC.HaveStack(ns_name) {
    err := l.TC.PositionStack(ns_name)
    if err != nil {
      l.SetError(fmt.Sprintf("Error entering NS: %v", err))
    }
  } else {
    AddNewStack(l, ns_name)
  }
}

func (l *TCExecListener) ExitNs(c *parser.NsContext) {
  if l.TC.Errors() > 0 {
    return
  }
  ns_name := c.GetNsname().GetText()
  log.Debugf("Exiting NS: %v", ns_name)
  l.TC.StacksLeft(1)
}
