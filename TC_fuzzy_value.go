package ThreadComputation

import (
  "strconv"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)

func (l *TCExecListener) EnterVal(c *parser.ValContext) {
  if c.GetFname() == nil {
    l.SetError("Can not get value out of context")
    return
  }
  func_name := c.GetFname().GetText()
  log.Debugf("Begin to create a Value[%v]", func_name)
  l.TC.InAttr += 1
  l.TC.Attrs.Add()
}

func (l *TCExecListener) ExitVal(c *parser.ValContext) {
  var p float64
  var err error
  if c.GetFname() == nil {
    l.SetError("Can not get value out of context")
    return
  }
  p = 100.0
  if c.GetP() != nil {
    p_text := c.GetP().GetText()
    p, err = strconv.ParseFloat(p_text, 64)
    if err != nil {
      p = 100.0
    }
  }
  q := l.Attrs()
  if q.Len() == 0 {
    log.Debugf("Nothing has been passed for a value creation")
    return
  }
  log.Debugf("Setting values with p=%v", p)
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      ReturnFromFunction(l, "#VALUE", MakeValue(e.(*TCValue).Value, float32((p + float64(e.(*TCValue).P)/2)), 0))
    default:
      ReturnFromFunction(l, "#VALUE", MakeValue(e, float32(p), 0))
    }
  }
}
