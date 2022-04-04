package ThreadComputation

import (

)



func TCConvertLinesToString(x interface{}) string {
  res := ""
  switch x.(type) {
  case *TCLines:
    res += x.(*TCLines).String()
  }
  return res
}

func init() {
  RegisterToStringCallback(Lines, TCConvertLinesToString)
}
