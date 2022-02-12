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

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterDblock_term is called when entering the dblock_term production.
	EnterDblock_term(c *Dblock_termContext)

	// EnterDblock is called when entering the dblock production.
	EnterDblock(c *DblockContext)

	// EnterInteger_term is called when entering the integer_term production.
	EnterInteger_term(c *Integer_termContext)

	// EnterFloat_term is called when entering the float_term production.
	EnterFloat_term(c *Float_termContext)

	// EnterString_term is called when entering the string_term production.
	EnterString_term(c *String_termContext)

	// EnterBoolean_term is called when entering the boolean_term production.
	EnterBoolean_term(c *Boolean_termContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitDblock_term is called when exiting the dblock_term production.
	ExitDblock_term(c *Dblock_termContext)

	// ExitDblock is called when exiting the dblock production.
	ExitDblock(c *DblockContext)

	// ExitInteger_term is called when exiting the integer_term production.
	ExitInteger_term(c *Integer_termContext)

	// ExitFloat_term is called when exiting the float_term production.
	ExitFloat_term(c *Float_termContext)

	// ExitString_term is called when exiting the string_term production.
	ExitString_term(c *String_termContext)

	// ExitBoolean_term is called when exiting the boolean_term production.
	ExitBoolean_term(c *Boolean_termContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)
}
