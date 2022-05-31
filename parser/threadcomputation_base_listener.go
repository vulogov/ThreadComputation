// Code generated from ThreadComputation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreadComputation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseThreadComputationListener is a complete listener for a parse tree produced by ThreadComputationParser.
type BaseThreadComputationListener struct{}

var _ ThreadComputationListener = &BaseThreadComputationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseThreadComputationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseThreadComputationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseThreadComputationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseThreadComputationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseThreadComputationListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseThreadComputationListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterRoot_term is called when production root_term is entered.
func (s *BaseThreadComputationListener) EnterRoot_term(ctx *Root_termContext) {}

// ExitRoot_term is called when production root_term is exited.
func (s *BaseThreadComputationListener) ExitRoot_term(ctx *Root_termContext) {}

// EnterNs is called when production ns is entered.
func (s *BaseThreadComputationListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BaseThreadComputationListener) ExitNs(ctx *NsContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseThreadComputationListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseThreadComputationListener) ExitFun(ctx *FunContext) {}

// EnterVal is called when production val is entered.
func (s *BaseThreadComputationListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BaseThreadComputationListener) ExitVal(ctx *ValContext) {}

// EnterNs_term is called when production ns_term is entered.
func (s *BaseThreadComputationListener) EnterNs_term(ctx *Ns_termContext) {}

// ExitNs_term is called when production ns_term is exited.
func (s *BaseThreadComputationListener) ExitNs_term(ctx *Ns_termContext) {}

// EnterFun_term is called when production fun_term is entered.
func (s *BaseThreadComputationListener) EnterFun_term(ctx *Fun_termContext) {}

// ExitFun_term is called when production fun_term is exited.
func (s *BaseThreadComputationListener) ExitFun_term(ctx *Fun_termContext) {}

// EnterPos_term is called when production pos_term is entered.
func (s *BaseThreadComputationListener) EnterPos_term(ctx *Pos_termContext) {}

// ExitPos_term is called when production pos_term is exited.
func (s *BaseThreadComputationListener) ExitPos_term(ctx *Pos_termContext) {}
