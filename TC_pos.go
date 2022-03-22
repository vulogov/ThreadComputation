package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)


func GetPositionalFromAttr(l *TCExecListener, n int64) interface{} {
  q := l.LastEvAttr()
  if q == nil {
    log.Debug("There is no attrs in evaluiation")
    return nil
  }
  if n < 0 && int(n) > q.Len() {
    log.Debugf("Postion %v is outside attr %v", n, q.Len())
    return nil
  }
  return q.At(int(n))
}


func (l *TCExecListener) ExitPos_term(c *parser.Pos_termContext) {
  if l.TC.Errors() > 0 {
    return
  }

  pos_name := c.GetPname().GetText()
  if l.TC.AddToUserFun(fmt.Sprintf("#%v", pos_name)) {
    log.Debugf("Recorded positional to userfun: #%v", pos_name)
    return
  }
  log.Debugf("Positional: %v", pos_name)
  if IsSimpleData(pos_name) {
    pos := GetSimpleData(pos_name)
    log.Debugf("Postional: %v %T", pos_name, pos)
    switch pos.(type) {
    case int64:
      log.Debug("Positional parameter is numeric index from current attribute")
      e := GetPositionalFromAttr(l, pos.(int64))
      if e == nil {
        l.SetError("Can not get positional attribute from evaluation: %v", pos)
        return
      }
      ReturnFromFunction(l, "#POSITIONAL", e)
    }
  }
}
