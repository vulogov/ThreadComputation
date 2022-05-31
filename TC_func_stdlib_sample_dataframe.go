package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func TCDataframeSample(x interface{}, y interface{}) interface{} {
  out := MakeNumbers()

  switch data :=  x.(type) {
  case *TCData:
    switch y.(type) {
    case string:
      c, err := data.D.NameToColumn(y.(string))
      if err != nil {
        log.Errorf("DataFrame error: %v", err)
        return nil
      }
      iterator := data.D.Series[c].ValuesIterator()
      for  {
        n, d, _ := iterator()
        if n == nil {
          break
        }
        out.Add(d)
      }
    }
  }
  return out
}

func init() {
  RegisterOperatorCallback("sample", Data, String, TCDataframeSample)
}
