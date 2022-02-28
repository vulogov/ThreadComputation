package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
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

func (l *TCExecListener) EnterSpawnblock(c *parser.SpawnblockContext) {
  log.Debug("Entering SPAWN")
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("spawn\\") == true {
    return
  }
  MakeTempFun(l)
}

func (l *TCExecListener) ExitSpawnblock(c *parser.SpawnblockContext) {
  var func_name string

  if l.TC.Errors() > 0 {
    return
  }

  if l.TC.AddToUserFun("\\") == true {
    return
  }

  log.Debug("Exiting SPAWN")
  if l.TC.UFStack.Len() > 0 {
    func_name = l.TC.UFNStack.Front().(string)
    code := l.TC.UFStack.PopFront().(string)
    l.TC.UFNStack.PopFront()
    log.Debugf("Spawning block: %v", func_name)
    l.TC.GoEval(code)
  }
}

func (l *TCExecListener) EnterSendblock(c *parser.SendblockContext) {
  log.Debug("Entering SEND")
  if l.TC.Errors() > 0 {
    return
  }

  if l.TC.AddToUserFun("send\\") == true {
    return
  }

  MakeTempFun(l)
}

func (l *TCExecListener) ExitSendblock(c *parser.SendblockContext) {
  var func_name string

  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("\\") == true {
    return
  }
  log.Debug("Exiting SEND")
  if l.TC.UFStack.Len() > 0 {
    func_name = l.TC.UFNStack.Front().(string)
    code := l.TC.UFStack.PopFront().(string)
    l.TC.UFNStack.PopFront()
    log.Debugf("Sending block: %v", func_name)
    l.TC.Eval(code)
  }
}

func (l *TCExecListener) EnterRecvblock(c *parser.RecvblockContext) {
  log.Debug("Entering RECV")
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("recv\\") == true {
    return
  }
  MakeTempFun(l)
}

func (l *TCExecListener) ExitRecvblock(c *parser.RecvblockContext) {
  var func_name string

  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("\\") == true {
    return
  }
  log.Debug("Exiting RECV")
  if l.TC.UFStack.Len() > 0 {
    func_name = l.TC.UFNStack.Front().(string)
    code := l.TC.UFStack.PopFront().(string)
    l.TC.UFNStack.PopFront()
    log.Debugf("Receiving block: %v", func_name)
    l.TC.Eval(code)
  }
}

func TCWaitFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  log.Debug("Waiting for all goroutines to finish")
  l.TC.Wg.Wait()
  return nil, nil
}

func TCSendFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    if q.Len() > 0 {
      sname := q.PopFront()
      switch sname.(type) {
      case string:
        tcSendDataToChannel(l, sname.(string), e)
      }
    } else {
      sname := l.TC.ResNames.Front().(string)
      tcSendDataToChannel(l, sname, e)
    }
  }
  return nil, nil
}

func TCRecvFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    sname := q.PopFront()
    switch sname.(type) {
    case string:
      e := tcRecvDataFromChannel(l, sname.(string))
      if e != nil {
        ReturnFromFunction(l, "recv", e)
      }
    }
  } else {
    sname := l.TC.ResNames.Front().(string)
    e := tcRecvDataFromChannel(l, sname)
    if e != nil {
      ReturnFromFunction(l, "recv", e)
    }
  }
  return nil, nil
}


func initStdlibGoRoutines() {
  SetFunction("wait", TCWaitFunction)
  SetFunction("send", TCSendFunction)
  SetFunction("recv", TCRecvFunction)
}
