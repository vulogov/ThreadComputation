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

// EnterUfun is called when production ufun is entered.
func (s *BaseThreadComputationListener) EnterUfun(ctx *UfunContext) {}

// ExitUfun is called when production ufun is exited.
func (s *BaseThreadComputationListener) ExitUfun(ctx *UfunContext) {}

// EnterVars is called when production vars is entered.
func (s *BaseThreadComputationListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseThreadComputationListener) ExitVars(ctx *VarsContext) {}

// EnterDblock_term is called when production dblock_term is entered.
func (s *BaseThreadComputationListener) EnterDblock_term(ctx *Dblock_termContext) {}

// ExitDblock_term is called when production dblock_term is exited.
func (s *BaseThreadComputationListener) ExitDblock_term(ctx *Dblock_termContext) {}

// EnterFun_term is called when production fun_term is entered.
func (s *BaseThreadComputationListener) EnterFun_term(ctx *Fun_termContext) {}

// ExitFun_term is called when production fun_term is exited.
func (s *BaseThreadComputationListener) ExitFun_term(ctx *Fun_termContext) {}

// EnterUfun_term is called when production ufun_term is entered.
func (s *BaseThreadComputationListener) EnterUfun_term(ctx *Ufun_termContext) {}

// ExitUfun_term is called when production ufun_term is exited.
func (s *BaseThreadComputationListener) ExitUfun_term(ctx *Ufun_termContext) {}

// EnterDblock is called when production dblock is entered.
func (s *BaseThreadComputationListener) EnterDblock(ctx *DblockContext) {}

// ExitDblock is called when production dblock is exited.
func (s *BaseThreadComputationListener) ExitDblock(ctx *DblockContext) {}

// EnterLblock is called when production lblock is entered.
func (s *BaseThreadComputationListener) EnterLblock(ctx *LblockContext) {}

// ExitLblock is called when production lblock is exited.
func (s *BaseThreadComputationListener) ExitLblock(ctx *LblockContext) {}

// EnterDmap is called when production dmap is entered.
func (s *BaseThreadComputationListener) EnterDmap(ctx *DmapContext) {}

// ExitDmap is called when production dmap is exited.
func (s *BaseThreadComputationListener) ExitDmap(ctx *DmapContext) {}

// EnterInteger_term is called when production integer_term is entered.
func (s *BaseThreadComputationListener) EnterInteger_term(ctx *Integer_termContext) {}

// ExitInteger_term is called when production integer_term is exited.
func (s *BaseThreadComputationListener) ExitInteger_term(ctx *Integer_termContext) {}

// EnterFloat_term is called when production float_term is entered.
func (s *BaseThreadComputationListener) EnterFloat_term(ctx *Float_termContext) {}

// ExitFloat_term is called when production float_term is exited.
func (s *BaseThreadComputationListener) ExitFloat_term(ctx *Float_termContext) {}

// EnterString_term is called when production string_term is entered.
func (s *BaseThreadComputationListener) EnterString_term(ctx *String_termContext) {}

// ExitString_term is called when production string_term is exited.
func (s *BaseThreadComputationListener) ExitString_term(ctx *String_termContext) {}

// EnterBoolean_term is called when production boolean_term is entered.
func (s *BaseThreadComputationListener) EnterBoolean_term(ctx *Boolean_termContext) {}

// ExitBoolean_term is called when production boolean_term is exited.
func (s *BaseThreadComputationListener) ExitBoolean_term(ctx *Boolean_termContext) {}

// EnterKey_term is called when production key_term is entered.
func (s *BaseThreadComputationListener) EnterKey_term(ctx *Key_termContext) {}

// ExitKey_term is called when production key_term is exited.
func (s *BaseThreadComputationListener) ExitKey_term(ctx *Key_termContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseThreadComputationListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseThreadComputationListener) ExitTerm(ctx *TermContext) {}
