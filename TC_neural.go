package ThreadComputation
import (
  "fmt"
  "time"
  log "github.com/sirupsen/logrus"
  "github.com/goml/gobrain"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
)

type TCNeural struct {
  Id          string
  Stamp       time.Time
  Patterns    [][][]float64
  NN          *gobrain.FeedForward
  Inputs      int
  Outputs     int
  Hidden      int
  Epoch       int
  Lr          float64
  Momentum    float64
  IsTrained   bool
}

func (n *TCNeural) String() string {
  return fmt.Sprintf("neural[ input=%v, hidden=%v, output=%v pattern(s)=%v trained=%v ]", n.Inputs, n.Hidden, n.Outputs, len(n.Patterns),n.IsTrained)
}

func IsNeural(x interface{}) bool {
  switch x.(type) {
  case *TCNeural:
    return true
  }
  return false
}

func MakeNeural() *TCNeural {
  res := new(TCNeural)
  res.Id    = uuid.NewString()
  res.Stamp = time.Now()
  res.IsTrained = false
  res.Epoch = 1000
  res.Lr    = 0.6
  res.Momentum = 0.4
  res.NN    = new(gobrain.FeedForward)
  return res
}

func TCNeuralAddTrainSet(n *TCNeural, p *TCPair) *TCNeural {
  switch p.X.(type) {
  case *TCNumbers:
    switch p.Y.(type) {
    case *TCNumbers:
      if n.Inputs == 0 {
        n.Inputs = len(p.X.(*TCNumbers).N)
      }
      if n.Outputs == 0 {
        n.Outputs = len(p.Y.(*TCNumbers).N)
      }
      if n.Hidden < n.Inputs {
        n.Hidden = n.Inputs
      }
      if n.Inputs != len(p.X.(*TCNumbers).N) || n.Outputs != len(p.Y.(*TCNumbers).N) {
        log.Debugf("Bad trainin set for %s", n.String())
        return nil
      }
      n.Patterns = append(n.Patterns, [][]float64{p.X.(*TCNumbers).N, p.Y.(*TCNumbers).N})
      return n
    }
  }
  return nil
}

func TCTrainNeuralFromPairs(l *TCExecListener, n *TCNeural, q *deque.Deque) error {
  l.ExecuteOperatorTyped("set", Neural, q)
  return nil
}

func TCNeuralTrain(x interface{}, y interface{}) {
  n := x.(*TCNeural)
  s := y.(*TCPair)
  res := TCNeuralAddTrainSet(n, s)
  if res == nil {
    log.Debugf("Attrmpt to add training set failed for %v", n.String())
  }
}

func TCNeuralUpdate(x interface{}, y interface{}) interface{} {
  n := x.(*TCNeural)
  u := y.(*TCNumbers)
  if len(u.N) != n.Inputs {
    return MakeNone()
  }
  log.Debugf("Update neural with %T %v", u.N, u.N)
  r := n.NN.Update(u.N)
  res := MakeNumbers()
  res.Set(r)
  return res
}

func TCNeuralFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeNeural()
  q.PushFront(res)
  TCTrainNeuralFromPairs(l, res, q)
  return res, nil
}

func TCNeuralValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeNeural()
  q.PushFront(res)
  TCTrainNeuralFromPairs(l, res, q)
  return MakeValue(res, 100.0, 0), nil
}

func TCTrainNeuralFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  rn := l.AttrByType(Neural, q)
  if rn == nil {
    return nil, l.TC.MakeError("neural[] not found for train[]")
  }
  n := rn.(*TCNeural)
  if n.IsTrained {
    log.Debugf("neural[%v] already trained", n.Id)
    return nil, nil
  }
  err := TCTrainNeuralFromPairs(l, n, q)
  if err != nil {
    return nil, err
  }
  n.NN.Init(n.Inputs, n.Hidden, n.Outputs)
  log.Debugf("E=%v LR=%v M=%v Train data: %v", n.Epoch, n.Lr, n.Momentum, n.Patterns)

  if l.TC.IsDebug {
    n.NN.Train(n.Patterns, n.Epoch, n.Lr, n.Momentum, true)
  } else {
    n.NN.Train(n.Patterns, n.Epoch, n.Lr, n.Momentum, false)
  }
  n.IsTrained = true
  return n, nil
}

func init() {
  SetCommand("neural", TCNeuralFunction)
  SetCommand("JustNeural", TCNeuralValueFunction)
  SetFunction("train", TCTrainNeuralFunction)
  RegisterOperatorCmdCallback("set", Neural, Pair, TCNeuralTrain)
  RegisterOperatorCallback("get", Neural, Numbers, TCNeuralUpdate)
}
