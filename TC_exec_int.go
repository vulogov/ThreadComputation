package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)

func (l *TCExecListener) EnterInteger_term(c *parser.Integer_termContext) {
  if l.TC.Errors() > 0 {
    return
  }
  str, err := conv.Int64(c.GetVALUE().GetText())
  if err == nil {
    if l.TC.InAttr < 1 {
      l.TC.Res.PushFront(str)
    } else {
      l.TC.Attrs.Set(str)
    }
  }
}
