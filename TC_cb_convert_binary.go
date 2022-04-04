package ThreadComputation

import (
  "fmt"
  "strings"
  "strconv"
  "github.com/deckarep/golang-set"
  conv "github.com/cstockton/go-conv"
  log "github.com/sirupsen/logrus"
)


func TCBinaryConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCBinary:
    switch to_type {
    case String:
      return e.String()
    case Lines:
      return MakeLinesFromValue(e.(*TCBinary))
    }
  }
  return nil
}


func init() {
  RegisterConvertCallback(Binary, TCBinaryConvert)
}
