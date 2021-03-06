// Code generated from ThreadComputation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ThreadComputationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var threadcomputationlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func threadcomputationlexerLexerInit() {
	staticData := &threadcomputationlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'['", "':'", "';;'", "']'", "'{'", "'}'", "'#'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "FUNC_NAME", "OPS", "NAME", "INTEGER",
		"DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "MOD", "BLOCK_COMMENT",
		"WhiteSpace", "NewLine",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "FUNC_NAME",
		"OPS", "NAME", "INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING",
		"OP", "MOD", "BLOCK_COMMENT", "WhiteSpace", "NewLine", "NON_ZERO_DIGIT",
		"DIGIT", "SIGN", "EXPONENT_OR_POINT_FLOAT", "POINT_FLOAT", "RN", "SHORT_STRING",
		"ID_START", "ID_CONTINUE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 243, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 78, 8, 7, 1, 8, 4, 8, 81, 8, 8, 11, 8, 12, 8, 82, 1, 9, 1, 9,
		5, 9, 87, 8, 9, 10, 9, 12, 9, 90, 9, 9, 1, 10, 3, 10, 93, 8, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 5, 11, 99, 8, 11, 10, 11, 12, 11, 102, 9, 11, 1, 11,
		4, 11, 105, 8, 11, 11, 11, 12, 11, 106, 3, 11, 109, 8, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5,
		16, 123, 8, 16, 10, 16, 12, 16, 126, 9, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 17, 4, 17, 134, 8, 17, 11, 17, 12, 17, 135, 1, 17, 1, 17, 1,
		18, 3, 18, 141, 8, 18, 1, 18, 1, 18, 3, 18, 145, 8, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 3, 22, 156, 8, 22, 1, 22,
		4, 22, 159, 8, 22, 11, 22, 12, 22, 160, 1, 22, 3, 22, 164, 8, 22, 1, 22,
		1, 22, 3, 22, 168, 8, 22, 1, 22, 4, 22, 171, 8, 22, 11, 22, 12, 22, 172,
		1, 22, 3, 22, 176, 8, 22, 1, 22, 3, 22, 179, 8, 22, 1, 23, 5, 23, 182,
		8, 23, 10, 23, 12, 23, 185, 9, 23, 1, 23, 1, 23, 4, 23, 189, 8, 23, 11,
		23, 12, 23, 190, 1, 23, 4, 23, 194, 8, 23, 11, 23, 12, 23, 195, 1, 23,
		3, 23, 199, 8, 23, 1, 24, 3, 24, 202, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 3, 25, 210, 8, 25, 1, 25, 5, 25, 213, 8, 25, 10, 25, 12,
		25, 216, 9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 223, 8, 25, 1,
		25, 5, 25, 226, 8, 25, 10, 25, 12, 25, 229, 9, 25, 1, 25, 3, 25, 232, 8,
		25, 1, 26, 3, 26, 235, 8, 26, 1, 26, 3, 26, 238, 8, 26, 1, 27, 1, 27, 3,
		27, 242, 8, 27, 1, 124, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 0, 51, 0,
		53, 0, 55, 0, 1, 0, 11, 15, 0, 33, 33, 38, 38, 42, 45, 47, 47, 59, 62,
		94, 95, 124, 124, 172, 172, 215, 215, 247, 247, 955, 955, 8593, 8593, 8595,
		8595, 8710, 8711, 8743, 8744, 8, 0, 36, 36, 63, 64, 96, 96, 126, 126, 8704,
		8704, 8707, 8708, 8728, 8728, 9633, 9633, 2, 0, 9, 9, 32, 32, 1, 0, 49,
		57, 1, 0, 48, 57, 2, 0, 43, 43, 45, 45, 2, 0, 69, 69, 101, 101, 4, 0, 10,
		10, 13, 13, 39, 39, 92, 92, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 2, 0,
		65, 90, 97, 122, 1, 0, 97, 122, 268, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5, 61,
		1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 66, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13,
		70, 1, 0, 0, 0, 15, 77, 1, 0, 0, 0, 17, 80, 1, 0, 0, 0, 19, 84, 1, 0, 0,
		0, 21, 92, 1, 0, 0, 0, 23, 108, 1, 0, 0, 0, 25, 110, 1, 0, 0, 0, 27, 112,
		1, 0, 0, 0, 29, 114, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 118, 1, 0, 0,
		0, 35, 133, 1, 0, 0, 0, 37, 144, 1, 0, 0, 0, 39, 148, 1, 0, 0, 0, 41, 150,
		1, 0, 0, 0, 43, 152, 1, 0, 0, 0, 45, 178, 1, 0, 0, 0, 47, 198, 1, 0, 0,
		0, 49, 201, 1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 237, 1, 0, 0, 0, 55, 241,
		1, 0, 0, 0, 57, 58, 5, 91, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60, 5, 58, 0, 0,
		60, 4, 1, 0, 0, 0, 61, 62, 5, 59, 0, 0, 62, 63, 5, 59, 0, 0, 63, 6, 1,
		0, 0, 0, 64, 65, 5, 93, 0, 0, 65, 8, 1, 0, 0, 0, 66, 67, 5, 123, 0, 0,
		67, 10, 1, 0, 0, 0, 68, 69, 5, 125, 0, 0, 69, 12, 1, 0, 0, 0, 70, 71, 5,
		35, 0, 0, 71, 14, 1, 0, 0, 0, 72, 78, 3, 19, 9, 0, 73, 78, 3, 17, 8, 0,
		74, 78, 3, 21, 10, 0, 75, 78, 3, 25, 12, 0, 76, 78, 3, 27, 13, 0, 77, 72,
		1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		77, 76, 1, 0, 0, 0, 78, 16, 1, 0, 0, 0, 79, 81, 3, 29, 14, 0, 80, 79, 1,
		0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83,
		18, 1, 0, 0, 0, 84, 88, 3, 53, 26, 0, 85, 87, 3, 55, 27, 0, 86, 85, 1,
		0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		20, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 93, 3, 43, 21, 0, 92, 91, 1, 0,
		0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 3, 23, 11, 0, 95,
		22, 1, 0, 0, 0, 96, 100, 3, 39, 19, 0, 97, 99, 3, 41, 20, 0, 98, 97, 1,
		0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0,
		101, 109, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 105, 5, 48, 0, 0, 104,
		103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 96, 1, 0, 0, 0, 108, 104, 1, 0,
		0, 0, 109, 24, 1, 0, 0, 0, 110, 111, 3, 45, 22, 0, 111, 26, 1, 0, 0, 0,
		112, 113, 3, 51, 25, 0, 113, 28, 1, 0, 0, 0, 114, 115, 7, 0, 0, 0, 115,
		30, 1, 0, 0, 0, 116, 117, 7, 1, 0, 0, 117, 32, 1, 0, 0, 0, 118, 119, 5,
		47, 0, 0, 119, 120, 5, 42, 0, 0, 120, 124, 1, 0, 0, 0, 121, 123, 9, 0,
		0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		124, 122, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127,
		128, 5, 42, 0, 0, 128, 129, 5, 47, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131,
		6, 16, 0, 0, 131, 34, 1, 0, 0, 0, 132, 134, 7, 2, 0, 0, 133, 132, 1, 0,
		0, 0, 134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 138, 6, 17, 0, 0, 138, 36, 1, 0, 0, 0, 139,
		141, 5, 13, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 145, 5, 10, 0, 0, 143, 145, 5, 13, 0, 0, 144, 140, 1,
		0, 0, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 6, 18, 0,
		0, 147, 38, 1, 0, 0, 0, 148, 149, 7, 3, 0, 0, 149, 40, 1, 0, 0, 0, 150,
		151, 7, 4, 0, 0, 151, 42, 1, 0, 0, 0, 152, 153, 7, 5, 0, 0, 153, 44, 1,
		0, 0, 0, 154, 156, 3, 43, 21, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 163, 1, 0, 0, 0, 157, 159, 7, 4, 0, 0, 158, 157, 1, 0, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		164, 1, 0, 0, 0, 162, 164, 3, 47, 23, 0, 163, 158, 1, 0, 0, 0, 163, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 167, 7, 6, 0, 0, 166, 168, 7, 5,
		0, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0,
		169, 171, 7, 4, 0, 0, 170, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 179, 1, 0, 0, 0, 174, 176,
		3, 43, 21, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 1,
		0, 0, 0, 177, 179, 3, 47, 23, 0, 178, 155, 1, 0, 0, 0, 178, 175, 1, 0,
		0, 0, 179, 46, 1, 0, 0, 0, 180, 182, 7, 4, 0, 0, 181, 180, 1, 0, 0, 0,
		182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188, 5, 46, 0, 0, 187, 189,
		7, 4, 0, 0, 188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 199, 1, 0, 0, 0, 192, 194, 7, 4, 0, 0,
		193, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 5, 46, 0, 0, 198, 183,
		1, 0, 0, 0, 198, 193, 1, 0, 0, 0, 199, 48, 1, 0, 0, 0, 200, 202, 5, 13,
		0, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0,
		203, 204, 5, 10, 0, 0, 204, 50, 1, 0, 0, 0, 205, 214, 5, 39, 0, 0, 206,
		209, 5, 92, 0, 0, 207, 210, 3, 49, 24, 0, 208, 210, 9, 0, 0, 0, 209, 207,
		1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 213, 8, 7,
		0, 0, 212, 206, 1, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0,
		214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216,
		214, 1, 0, 0, 0, 217, 232, 5, 39, 0, 0, 218, 227, 5, 34, 0, 0, 219, 222,
		5, 92, 0, 0, 220, 223, 3, 49, 24, 0, 221, 223, 9, 0, 0, 0, 222, 220, 1,
		0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 226, 8, 8, 0,
		0, 225, 219, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229, 227,
		1, 0, 0, 0, 230, 232, 5, 34, 0, 0, 231, 205, 1, 0, 0, 0, 231, 218, 1, 0,
		0, 0, 232, 52, 1, 0, 0, 0, 233, 235, 7, 9, 0, 0, 234, 233, 1, 0, 0, 0,
		235, 238, 1, 0, 0, 0, 236, 238, 7, 10, 0, 0, 237, 234, 1, 0, 0, 0, 237,
		236, 1, 0, 0, 0, 238, 54, 1, 0, 0, 0, 239, 242, 3, 53, 26, 0, 240, 242,
		5, 46, 0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 56, 1, 0,
		0, 0, 34, 0, 77, 82, 88, 92, 100, 106, 108, 124, 135, 140, 144, 155, 160,
		163, 167, 172, 175, 178, 183, 190, 195, 198, 201, 209, 212, 214, 222, 225,
		227, 231, 234, 237, 241, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ThreadComputationLexerInit initializes any static state used to implement ThreadComputationLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewThreadComputationLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ThreadComputationLexerInit() {
	staticData := &threadcomputationlexerLexerStaticData
	staticData.once.Do(threadcomputationlexerLexerInit)
}

// NewThreadComputationLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewThreadComputationLexer(input antlr.CharStream) *ThreadComputationLexer {
	ThreadComputationLexerInit()
	l := new(ThreadComputationLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &threadcomputationlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ThreadComputation.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ThreadComputationLexer tokens.
const (
	ThreadComputationLexerT__0            = 1
	ThreadComputationLexerT__1            = 2
	ThreadComputationLexerT__2            = 3
	ThreadComputationLexerT__3            = 4
	ThreadComputationLexerT__4            = 5
	ThreadComputationLexerT__5            = 6
	ThreadComputationLexerT__6            = 7
	ThreadComputationLexerFUNC_NAME       = 8
	ThreadComputationLexerOPS             = 9
	ThreadComputationLexerNAME            = 10
	ThreadComputationLexerINTEGER         = 11
	ThreadComputationLexerDECIMAL_INTEGER = 12
	ThreadComputationLexerFLOAT_NUMBER    = 13
	ThreadComputationLexerSTRING          = 14
	ThreadComputationLexerOP              = 15
	ThreadComputationLexerMOD             = 16
	ThreadComputationLexerBLOCK_COMMENT   = 17
	ThreadComputationLexerWhiteSpace      = 18
	ThreadComputationLexerNewLine         = 19
)
