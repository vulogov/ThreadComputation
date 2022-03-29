package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/hbollon/go-edlib"
)





func TCCompareStringString(name string, x interface{}, y interface{}) interface{} {
  res, err := edlib.StringsSimilarity(x.(string), y.(string), edlib.Levenshtein)
  if err != nil {
    return nil
  }
  log.Debugf("Distance between strings %v %v = %v", x, y, res)
  switch name {
  case "=":
    return res == 1.0
  case "!=":
    return res != 1.0
  case ">":
    return res > 0.5
  case "<":
    return res < 0.5
  case ">=":
    return res >= 0.5
  case "<=":
    return res <= 0.5
  }
  return nil
}

func init() {
  RegisterCompareCallback(String, String, TCCompareStringString)
}
