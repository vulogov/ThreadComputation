package ThreadComputation

import (
  "math"
  log "github.com/sirupsen/logrus"
)

func tccmpMatrixes(name string, x *TCMatrix, y *TCMatrix) bool {
  m1,n1 := x.M.Dims()
  m2,n2 := y.M.Dims()
  if m1 != m2 || n1 != n2 {
    log.Debugf("Matrixes are not the same arity %vx%v != %vx%x", m1, n1, m2, n2)
    return false
  }
  for i := 0; i < m1; i++ {
    for j := 0; j < n1; j++ {
      if math.IsNaN(x.M.At(i,j)) || math.IsNaN(y.M.At(i,j)) {
        continue
      }
      switch name {
      case "=":
        if x.M.At(i,j) != y.M.At(i,j) {
          return false
        }
      case "!=":
        if x.M.At(i,j) == y.M.At(i,j) {
          return false
        }
      case "<":
        if x.M.At(i,j) > y.M.At(i,j) {
          return false
        }
      case ">":
        if x.M.At(i,j) < y.M.At(i,j) {
          return false
        }
      case ">=":
        if x.M.At(i,j) <= y.M.At(i,j) {
          return false
        }
      case "<=":
        if x.M.At(i,j) >= y.M.At(i,j) {
          return false
        }
      default:
        return false
      }
    }
  }
  return true
}

func TCCompareMatrixMatrix(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  case "!=":
    return ! tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  case "<":
    return tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  case ">":
    return ! tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  case ">=":
    return tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  case "<=":
    return ! tccmpMatrixes(name, x.(*TCMatrix), y.(*TCMatrix))
  }
  return nil
}

func init() {
  RegisterCompareCallback(Matrix, Matrix, TCCompareMatrixMatrix)
}
