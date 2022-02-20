package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/google/uuid"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterDblock(c *parser.DblockContext) {
  var bname string
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("(") == true {
    return
  }
  b := c.GetBname()
  if b == nil {
    bname = uuid.NewString()
  } else {
    bname = b.GetText()
  }
  AddNewStack(l, bname)
}

func (l *TCExecListener) ExitDblock(c *parser.DblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun(")") == true {
    return
  }
  l.TC.Res.Left()
}

func AddNewStack(l *TCExecListener, name string) {
  l.TC.AddNewStack(name)
}

func (tc *TCstate) AddNewStack(name string) {
  log.Debugf("Creating new stack: %v", name)
  if tc.ResN.Contains(name) {
    log.Debugf("Can not create new stack as it is already exists: %v", name)
    return
  }
  tc.ResN.Add(name)
  tc.ResNames.PushFront(name)
  tc.Res.Add()
}

func (tc *TCstate) DropLastStack() error {
  if tc.ResNames.Len() == 0 {
    return errors.New("TwoStack is empty")
  }
  name := tc.ResNames.PopFront()
  log.Debugf("Dropping stack: %v", name)
  if tc.ResN.Contains(name) {
    tc.ResN.Remove(name)
  }
  tc.Res.Del()
  return nil
}

func (tc *TCstate) RotateStackNames() {
  n := tc.ResNames.PopFront()
  log.Debugf("Positioning stack through: %v", n)
  tc.ResNames.PushBack(n)
}

func (tc *TCstate) PositionStack(name string) error {
  if tc.ResNames.Len() == 0 {
    return errors.New("TwoStack is empty")
  }
  if tc.ResN.Contains(name) == false {
    return errors.New(fmt.Sprintf("Stack not found for postioning: %v", name))
  }
  for {
    if tc.ResNames.Front().(string) == name {
      log.Debugf("Stack positioned: %v", name)
      return nil
    }
    tc.RotateStackNames()
  }
}
