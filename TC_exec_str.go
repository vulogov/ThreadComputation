package ThreadComputation

import (
  "strings"
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)

func (l *TCExecListener) EnterString_term(c *parser.String_termContext) {
  if l.TC.Errors() > 0 {
    return
  }
  str, err := conv.String(c.GetVALUE().GetText())
  if l.TC.AddToUserFun(str) {
    return
  }
  if strings.HasPrefix(str, "\"") {
    str = strings.TrimSuffix(str, "\"")
    str = strings.TrimPrefix(str, "\"")
  }
  if strings.HasPrefix(str, "'") {
    str = strings.TrimSuffix(str, "'")
    str = strings.TrimPrefix(str, "'")
  }
  if err == nil {
    if l.TC.InAttr < 1 {
      l.TC.Res.Set(str)
    } else {
      l.TC.Attrs.Set(str)
    }
  }
}
