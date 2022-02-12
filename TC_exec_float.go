package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)

func (l *TCExecListener) EnterFloat_term(c *parser.Float_termContext) {
  if l.TC.Errors() > 0 {
    return
  }
  str, err := conv.Float64(c.GetVALUE().GetText())
  if err == nil {
    if l.TC.InAttr < 1 {
      l.TC.Res.Set(str)
    } else {
      l.TC.Attrs.Set(str)
    }
  }
}
