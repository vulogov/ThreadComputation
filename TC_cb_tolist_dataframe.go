package ThreadComputation

import (
  "github.com/rocketlaunchr/dataframe-go"
)



func TCConvertDataframeToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCData:
    iterator := e.D.ValuesIterator(dataframe.ValuesOptions{0,1,true})
    res := new(TCList)
    e.D.Lock()
    for {
      row, vals, _ := iterator()
      d := MakeDict()
      if row == nil {
        break
      }
      for key, value := range vals {
        d.D.Set(key, value)
      }
      res.Add(d)
    }
    e.D.Unlock()
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(Data, TCConvertDataframeToList)
}
