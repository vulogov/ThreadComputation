package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func TCRotateNumbersLeft(v interface{}) interface{} {
  log.Debugf("Rotate numbers left: %T", v)
  switch l := v.(type) {
  case *TCNumbers:
    l.RotateLeft()
    return l
  }
  return v
}

func TCRotateNumbersRight(v interface{}) interface{} {
  log.Debugf("Rotate numbers right: %T", v)
  switch l := v.(type) {
  case *TCNumbers:
    l.RotateRight()
    return l
  }
  return v
}

func init() {
  RegisterFunctionCallback("rotateright", Numbers, TCRotateNumbersRight)
  RegisterFunctionCallback("rotateleft", Numbers, TCRotateNumbersLeft)
}
