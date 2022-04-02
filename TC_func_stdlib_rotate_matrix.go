package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func TCRotateMatrixLeft(v interface{}) interface{} {
  log.Debugf("Rotate matrix left: %T", v)
  switch l := v.(type) {
  case *TCMatrix:
    return l.RotateLeft()
  }
  return v
}

func TCRotateMatrixRight(v interface{}) interface{} {
  log.Debugf("Rotate matrix right: %T", v)
  switch l := v.(type) {
  case *TCMatrix:
    return l.RotateRight()
  }
  return v
}

func init() {
  RegisterFunctionCallback("rotateright", Matrix, TCRotateMatrixRight)
  RegisterFunctionCallback("rotateleft", Matrix, TCRotateMatrixLeft)
}
