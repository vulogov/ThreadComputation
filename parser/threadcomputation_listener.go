// Code generated from ThreadComputation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreadComputation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ThreadComputationListener is a complete listener for a parse tree produced by ThreadComputationParser.
type ThreadComputationListener interface {
	antlr.ParseTreeListener

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterRoot_term is called when entering the root_term production.
	EnterRoot_term(c *Root_termContext)

	// EnterNs is called when entering the ns production.
	EnterNs(c *NsContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// EnterNs_term is called when entering the ns_term production.
	EnterNs_term(c *Ns_termContext)

	// EnterFun_term is called when entering the fun_term production.
	EnterFun_term(c *Fun_termContext)

	// EnterPos_term is called when entering the pos_term production.
	EnterPos_term(c *Pos_termContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)

	// ExitNs_term is called when exiting the ns_term production.
	ExitNs_term(c *Ns_termContext)

	// ExitFun_term is called when exiting the fun_term production.
	ExitFun_term(c *Fun_termContext)

	// ExitPos_term is called when exiting the pos_term production.
	ExitPos_term(c *Pos_termContext)
}
