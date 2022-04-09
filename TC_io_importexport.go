package ThreadComputation
//
// import (
//   "fmt"
//   "path"
//   "github.com/srfrog/dict"
//   "github.com/gammazero/deque"
// )
//
//
//
// func tcImportFromIO(io *TCIO) {
//   for item := range io.D.Items() {
//     path := item.Value
//     fun := TCIOGetReadRouter(path.(string))
//     if fun != nil {
//       res := fun(io, path.(string))
//       if res != nil {
//         exp := MakeExported(res)
//         expv := exp.Export()
//         if expv != nil {
//           ReturnFromFunction(io.L, "Read", expv)
//         }
//       }
//     } else {
//       res, err := ReadFile(path.(string))
//       if err != nil {
//         io.L.SetError(fmt.Sprintf("IO error: %v", err))
//       } else {
//         exp := MakeExported(res)
//         expv := exp.Export()
//         if expv != nil {
//           ReturnFromFunction(io.L, "Read", expv)
//         }
//       }
//     }
//   }
// }
//
// func TCImportFromIO(v interface{}) interface{} {
//   switch io := v.(type) {
//   case *TCIO:
//     tcImportFromIO(io)
//   }
//   return nil
// }
//
// func TCIOImportFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
//   err := l.ExecuteSingleArgumentFunction("ioimport", q)
//   if err != nil {
//     return nil, err
//   }
//   return nil, nil
// }
//
// func tcExportToIO(io *TCIO, d interface{}) {
//   for item := range io.D.Items() {
//     path := item.Value
//     fun := TCIOGetWriteRouter(path.(string))
//     if fun != nil {
//       res := fun(io, path.(string), d)
//       if res != nil {
//         ReturnFromFunction(io.L, "Export", res)
//       }
//     } else {
//       switch v := d.(type) {
//       case *TCBinary:
//         exp := MakeExportedFromBinary(v.(*TCBinary))
//         err := writeVfsFile(path.(string), exp.Data)
//         if err != nil {
//           io.L.SetError(fmt.Sprintf("IO error: %v", err))
//           return
//         }
//       case *TCExported:
//         err := writeVfsFile(path.(string), v.Data)
//         if err != nil {
//           io.L.SetError(fmt.Sprintf("IO error: %v", err))
//           return
//         }
//       }
//     }
//   }
// }
//
// func TCExportToIO(x interface{}, y interface{}) {
//   switch d := x.(type) {
//   case *TCIO:
//     tcExportToIO(d, y)
//   }
// }
//
// func TCIOExportFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
//   err := l.ExecuteOperatorCommand("ioexport", q)
//   if err != nil {
//     return nil, err
//   }
//   return nil, nil
// }
//
// func init() {
//   SetFunction("Import", TCIOImportFunction)
//   SetFunction("Export", TCIOExportFunction)
//   RegisterFunctionCallback("ioimport", IO, TCImportFromIO)
//   RegisterOperatorCmdCallback("ioexport", IO, Binary, TCExportToIO)
//   RegisterOperatorCmdCallback("ioexport", IO, Exported, TCExportToIO)
// }
