package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func TCExitCbPoolClose(l *TCExecListener) {
  log.Debugf("Closing workers pool. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
  l.TC.Pool.Close()
}

func TCExitCbSetExitError(l *TCExecListener) {
  log.Debugf("Set an error state for TC")
  l.TC.MakeError("[ BUND ] instance has been terminated")
}
