package ThreadComputation
import (
  "fmt"
  "time"
  log "github.com/sirupsen/logrus"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
  "github.com/rocketlaunchr/dataframe-go"
)

type TCData struct {
  Id          string
  Stamp       time.Time
  D          *dataframe.DataFrame
}

func (d *TCData) String() string {
  return fmt.Sprintf(d.D.Table())
}

func IsData(x interface{}) bool {
  switch x.(type) {
  case *TCData:
    return true
  }
  return false
}

func MakeData() *TCData {
  res := new(TCData)
  res.Id    = uuid.NewString()
  res.Stamp = time.Now()
  res.D     = dataframe.NewDataFrame()
  return res
}

func TCDataAddSeries(d *TCData, p *TCPair) error {
  var name   string
  var x_type int
  var s      dataframe.Series
  switch p.X.(type) {
  case string:
    name = p.X.(string)
  default:
    return MakeError("Expecting series name")
  }
  switch p.Y.(type) {
  case int64:
    x_type = int(p.Y.(int64))
  case *Type:
    x_type = p.Y.(*Type).T
  default:
    return MakeError("Expecting series type")
  }
  names := d.D.Names()
  if IsStringElementExist(names, name) {
    d.D.RemoveSeries(name)
  }
  switch x_type {
  case Int:
    ls := dataframe.NewSeriesInt64(name, nil)
    s = ls.NewSeries(name, nil)
  case Float:
    ls := dataframe.NewSeriesFloat64(name, nil)
    s = ls.NewSeries(name, nil)
  case String:
    ls := dataframe.NewSeriesString(name, nil)
    s = ls.NewSeries(name, nil)
  case Time:
    ls := dataframe.NewSeriesTime(name, nil)
    s = ls.NewSeries(name, nil)
  default:
    ls := dataframe.NewSeriesMixed(name, nil)
    s = ls.NewSeries(name, nil)
  }
  d.D.AddSeries(s, nil)
  return nil
}

func TCDataSet(x interface{}, y interface{}) {
  n := x.(*TCData)
  s := y.(*TCPair)
  err := TCDataAddSeries(n, s)
  if err != nil {
    log.Debugf("Error adding set to data %v", err)
  }
}

func TCSetDataFromPairs(l *TCExecListener, d *TCData, q *deque.Deque) error {
  l.ExecuteOperatorTyped("set", Data, q)
  return nil
}

func TCDataFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeData()
  q.PushFront(res)
  TCSetDataFromPairs(l, res, q)
  return res, nil
}

func TCDataValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeData()
  q.PushFront(res)
  TCSetDataFromPairs(l, res, q)
  return MakeValue(res, 100.0, 0), nil
}

func init() {
  SetCommand("data", TCDataFunction)
  SetCommand("JustData", TCDataValueFunction)
  RegisterOperatorCmdCallback("set", Data, Pair, TCDataSet)
}
