// Code generated from ThreadComputation.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

// EnterFun is called when production fun is entered.
func (s *BaseThreadComputationListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseThreadComputationListener) ExitFun(ctx *FunContext) {}

// EnterDblock_term is called when production dblock_term is entered.
func (s *BaseThreadComputationListener) EnterDblock_term(ctx *Dblock_termContext) {}

// ExitDblock_term is called when production dblock_term is exited.
func (s *BaseThreadComputationListener) ExitDblock_term(ctx *Dblock_termContext) {}

// EnterFun_term is called when production fun_term is entered.
func (s *BaseThreadComputationListener) EnterFun_term(ctx *Fun_termContext) {}

// ExitFun_term is called when production fun_term is exited.
func (s *BaseThreadComputationListener) ExitFun_term(ctx *Fun_termContext) {}

// EnterDblock is called when production dblock is entered.
func (s *BaseThreadComputationListener) EnterDblock(ctx *DblockContext) {}

// ExitDblock is called when production dblock is exited.
func (s *BaseThreadComputationListener) ExitDblock(ctx *DblockContext) {}

// EnterDmap is called when production dmap is entered.
func (s *BaseThreadComputationListener) EnterDmap(ctx *DmapContext) {}

// ExitDmap is called when production dmap is exited.
func (s *BaseThreadComputationListener) ExitDmap(ctx *DmapContext) {}

// EnterKey_term is called when production key_term is entered.
func (s *BaseThreadComputationListener) EnterKey_term(ctx *Key_termContext) {}

// ExitKey_term is called when production key_term is exited.
func (s *BaseThreadComputationListener) ExitKey_term(ctx *Key_termContext) {}
