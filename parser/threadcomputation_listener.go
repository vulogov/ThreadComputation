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

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterUfun is called when entering the ufun production.
	EnterUfun(c *UfunContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterDblock_term is called when entering the dblock_term production.
	EnterDblock_term(c *Dblock_termContext)

	// EnterFun_term is called when entering the fun_term production.
	EnterFun_term(c *Fun_termContext)

	// EnterUfun_term is called when entering the ufun_term production.
	EnterUfun_term(c *Ufun_termContext)

	// EnterDblock is called when entering the dblock production.
	EnterDblock(c *DblockContext)

	// EnterLblock is called when entering the lblock production.
	EnterLblock(c *LblockContext)

	// EnterTrueblock is called when entering the trueblock production.
	EnterTrueblock(c *TrueblockContext)

	// EnterFalseblock is called when entering the falseblock production.
	EnterFalseblock(c *FalseblockContext)

	// EnterFilterblock is called when entering the filterblock production.
	EnterFilterblock(c *FilterblockContext)

	// EnterSpawnblock is called when entering the spawnblock production.
	EnterSpawnblock(c *SpawnblockContext)

	// EnterSendblock is called when entering the sendblock production.
	EnterSendblock(c *SendblockContext)

	// EnterRecvblock is called when entering the recvblock production.
	EnterRecvblock(c *RecvblockContext)

	// EnterDmap is called when entering the dmap production.
	EnterDmap(c *DmapContext)

	// EnterInteger_term is called when entering the integer_term production.
	EnterInteger_term(c *Integer_termContext)

	// EnterFloat_term is called when entering the float_term production.
	EnterFloat_term(c *Float_termContext)

	// EnterString_term is called when entering the string_term production.
	EnterString_term(c *String_termContext)

	// EnterBoolean_term is called when entering the boolean_term production.
	EnterBoolean_term(c *Boolean_termContext)

	// EnterKey_term is called when entering the key_term production.
	EnterKey_term(c *Key_termContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitUfun is called when exiting the ufun production.
	ExitUfun(c *UfunContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitDblock_term is called when exiting the dblock_term production.
	ExitDblock_term(c *Dblock_termContext)

	// ExitFun_term is called when exiting the fun_term production.
	ExitFun_term(c *Fun_termContext)

	// ExitUfun_term is called when exiting the ufun_term production.
	ExitUfun_term(c *Ufun_termContext)

	// ExitDblock is called when exiting the dblock production.
	ExitDblock(c *DblockContext)

	// ExitLblock is called when exiting the lblock production.
	ExitLblock(c *LblockContext)

	// ExitTrueblock is called when exiting the trueblock production.
	ExitTrueblock(c *TrueblockContext)

	// ExitFalseblock is called when exiting the falseblock production.
	ExitFalseblock(c *FalseblockContext)

	// ExitFilterblock is called when exiting the filterblock production.
	ExitFilterblock(c *FilterblockContext)

	// ExitSpawnblock is called when exiting the spawnblock production.
	ExitSpawnblock(c *SpawnblockContext)

	// ExitSendblock is called when exiting the sendblock production.
	ExitSendblock(c *SendblockContext)

	// ExitRecvblock is called when exiting the recvblock production.
	ExitRecvblock(c *RecvblockContext)

	// ExitDmap is called when exiting the dmap production.
	ExitDmap(c *DmapContext)

	// ExitInteger_term is called when exiting the integer_term production.
	ExitInteger_term(c *Integer_termContext)

	// ExitFloat_term is called when exiting the float_term production.
	ExitFloat_term(c *Float_termContext)

	// ExitString_term is called when exiting the string_term production.
	ExitString_term(c *String_termContext)

	// ExitBoolean_term is called when exiting the boolean_term production.
	ExitBoolean_term(c *Boolean_termContext)

	// ExitKey_term is called when exiting the key_term production.
	ExitKey_term(c *Key_termContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
