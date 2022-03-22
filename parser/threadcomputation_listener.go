// Code generated from ThreadComputation.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

	// EnterNs_term is called when entering the ns_term production.
	EnterNs_term(c *Ns_termContext)

	// EnterDblock_term is called when entering the dblock_term production.
	EnterDblock_term(c *Dblock_termContext)

	// EnterFun_term is called when entering the fun_term production.
	EnterFun_term(c *Fun_termContext)

	// EnterPos_term is called when entering the pos_term production.
	EnterPos_term(c *Pos_termContext)

	// EnterDblock is called when entering the dblock production.
	EnterDblock(c *DblockContext)

	// EnterDmap is called when entering the dmap production.
	EnterDmap(c *DmapContext)

	// EnterKey_term is called when entering the key_term production.
	EnterKey_term(c *Key_termContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitNs_term is called when exiting the ns_term production.
	ExitNs_term(c *Ns_termContext)

	// ExitDblock_term is called when exiting the dblock_term production.
	ExitDblock_term(c *Dblock_termContext)

	// ExitFun_term is called when exiting the fun_term production.
	ExitFun_term(c *Fun_termContext)

	// ExitPos_term is called when exiting the pos_term production.
	ExitPos_term(c *Pos_termContext)

	// ExitDblock is called when exiting the dblock production.
	ExitDblock(c *DblockContext)

	// ExitDmap is called when exiting the dmap production.
	ExitDmap(c *DmapContext)

	// ExitKey_term is called when exiting the key_term production.
	ExitKey_term(c *Key_termContext)
}
