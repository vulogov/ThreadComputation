package ThreadComputation

import (
  "strings"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func tcSendDataToChannel(l *TCExecListener, name string, data interface{}) {
  if ch, ok := l.TC.StackChan.Load(name); ok {
    ch.(chan interface{}) <- data
    log.Debugf("sending %T to %v len()=%v", data, name, len(ch.(chan interface{})))
  }
}

func tcRecvDataFromChannel(l *TCExecListener, name string) interface{} {
  if ch, ok := l.TC.StackChan.Load(name); ok {
    if len(ch.(chan interface{})) > 0 {
      e := <- ch.(chan interface{})
      log.Debugf("receiving %T from %v len()=%v", e, name, len(ch.(chan interface{})))
      return e
    } else {
      log.Debugf("channel %v is empty", name)
    }
  }
  return nil
}

func TCWaitFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debug("Waiting for all goroutines to finish")
  l.TC.Wg.Wait()
  return nil, nil
}

func TCSendLocalFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.ResNames.Len() == 0 {
    return nil, l.TC.MakeError("Namespace context is too shallow for sending")
  }
  sname := l.TC.ResNames.Front().(string)
  for q.Len() > 0 {
    e := q.PopFront()
    tcSendDataToChannel(l, sname, e)
  }
  return nil, nil
}

func TCSendNamedFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, l.TC.MakeError("Context is too shallow for sending")
  }
  s := q.PopFront()
  switch s.(type) {
  case string:
    s = strings.ToLower(s.(string))
  default:
    return nil, l.TC.MakeError("Channel name is not defined")
  }
  for q.Len() > 0 {
    e := q.PopFront()
    tcSendDataToChannel(l, s.(string), e)
  }
  return nil, nil
}

func TCRecvLocalFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.ResNames.Len() == 0 {
    return nil, l.TC.MakeError("Namespace context is too shallow for receiving")
  }
  sname := l.TC.ResNames.Front().(string)
  for {
    e := tcRecvDataFromChannel(l, sname)
    if e == nil {
      break
    }
    ReturnFromFunction(l, "recv", e)
  }
  return nil, nil
}

func TCRecvNamedFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var sname string
  if l.TC.ResNames.Len() == 0 {
    return nil, l.TC.MakeError("Namespace context is too shallow for receiving")
  }
  if q.Len() == 0 {
    sname = l.TC.ResNames.Front().(string)
  } else {
    e := q.PopFront()
    switch e.(type) {
    case string:
      sname = e.(string)
    default:
      return nil, l.TC.MakeError("Bad arguments for receiving")
    }
  }
  out:
  for {
    e := tcRecvDataFromChannel(l, sname)
    if e == nil {
      break
    }
    ReturnFromFunction(l, "recv", e)
    if q.Len() == 0 {
      break out
    }
    e = q.PopFront()
    switch e.(type) {
    case string:
      sname = e.(string)
    default:
      return nil, l.TC.MakeError("Bad arguments for receiving")
    }
  }
  return nil, nil
}

func BlockSpawn(l *TCExecListener, name string, code string) interface{} {
  log.Debugf("Spawning code in goroutine: %v", name)
  l.TC.GoEval(code)
  return nil
}

func init() {
  RegisterBlockCallback("spawn", BlockSpawn)
  SetCommand("wait", TCWaitFunction)
  SetFunction("send", TCSendLocalFunction)
  SetFunction("Send", TCSendNamedFunction)
  SetCommand("recv", TCRecvLocalFunction)
  SetFunction("Recv", TCRecvNamedFunction)
}
