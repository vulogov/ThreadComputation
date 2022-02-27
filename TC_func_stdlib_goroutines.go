package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

func (l *TCExecListener) EnterSpawnblock(c *parser.SpawnblockContext) {
  log.Debug("Entering SPAWN")
  if l.TC.Errors() > 0 {
    return
  }
  MakeTempFun(l)
}

func (l *TCExecListener) ExitSpawnblock(c *parser.SpawnblockContext) {
  var func_name string

  if l.TC.Errors() > 0 {
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

func TCWaitFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  log.Debug("Waiting for all goroutines to finish")
  l.TC.Wg.Wait()
  return nil, nil
}


func initStdlibGoRoutines() {
  SetFunction("wait", TCWaitFunction)
}
