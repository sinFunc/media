// Code generated from nmd.g4 by ANTLR 4.9.2. DO NOT EDIT.

package nmd

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 138,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 7, 15, 85, 10, 15, 12, 15, 14, 15, 88, 11, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 96, 10, 17, 13, 17, 14,
	17, 97, 3, 18, 6, 18, 101, 10, 18, 13, 18, 14, 18, 102, 3, 18, 3, 18, 7,
	18, 107, 10, 18, 12, 18, 14, 18, 110, 11, 18, 3, 18, 3, 18, 6, 18, 114,
	10, 18, 13, 18, 14, 18, 115, 5, 18, 118, 10, 18, 3, 19, 6, 19, 121, 10,
	19, 13, 19, 14, 19, 122, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 7, 22, 134, 10, 22, 12, 22, 14, 22, 137, 11, 22, 3, 86, 2,
	23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 2, 33, 17, 35, 18, 37, 19, 39, 2, 41,
	2, 43, 20, 3, 2, 5, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 50, 59, 5, 2, 67,
	92, 97, 97, 99, 124, 2, 144, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 50, 3, 2, 2, 2, 9, 54, 3,
	2, 2, 2, 11, 58, 3, 2, 2, 2, 13, 65, 3, 2, 2, 2, 15, 67, 3, 2, 2, 2, 17,
	69, 3, 2, 2, 2, 19, 71, 3, 2, 2, 2, 21, 73, 3, 2, 2, 2, 23, 75, 3, 2, 2,
	2, 25, 77, 3, 2, 2, 2, 27, 79, 3, 2, 2, 2, 29, 81, 3, 2, 2, 2, 31, 91,
	3, 2, 2, 2, 33, 95, 3, 2, 2, 2, 35, 117, 3, 2, 2, 2, 37, 120, 3, 2, 2,
	2, 39, 126, 3, 2, 2, 2, 41, 128, 3, 2, 2, 2, 43, 130, 3, 2, 2, 2, 45, 46,
	7, 61, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 47, 2, 2, 48, 49, 7, 64, 2,
	2, 49, 6, 3, 2, 2, 2, 50, 51, 7, 62, 2, 2, 51, 52, 7, 47, 2, 2, 52, 53,
	7, 64, 2, 2, 53, 8, 3, 2, 2, 2, 54, 55, 7, 62, 2, 2, 55, 56, 7, 47, 2,
	2, 56, 57, 7, 47, 2, 2, 57, 10, 3, 2, 2, 2, 58, 59, 7, 62, 2, 2, 59, 60,
	7, 47, 2, 2, 60, 61, 7, 101, 2, 2, 61, 62, 7, 106, 2, 2, 62, 63, 7, 99,
	2, 2, 63, 64, 7, 112, 2, 2, 64, 12, 3, 2, 2, 2, 65, 66, 7, 125, 2, 2, 66,
	14, 3, 2, 2, 2, 67, 68, 7, 46, 2, 2, 68, 16, 3, 2, 2, 2, 69, 70, 7, 127,
	2, 2, 70, 18, 3, 2, 2, 2, 71, 72, 7, 93, 2, 2, 72, 20, 3, 2, 2, 2, 73,
	74, 7, 95, 2, 2, 74, 22, 3, 2, 2, 2, 75, 76, 7, 66, 2, 2, 76, 24, 3, 2,
	2, 2, 77, 78, 7, 60, 2, 2, 78, 26, 3, 2, 2, 2, 79, 80, 7, 63, 2, 2, 80,
	28, 3, 2, 2, 2, 81, 86, 7, 41, 2, 2, 82, 85, 5, 31, 16, 2, 83, 85, 11,
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 89, 90, 7, 41, 2, 2, 90, 30, 3, 2, 2, 2, 91, 92, 7, 94, 2, 2, 92, 93,
	7, 41, 2, 2, 93, 32, 3, 2, 2, 2, 94, 96, 5, 39, 20, 2, 95, 94, 3, 2, 2,
	2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 34,
	3, 2, 2, 2, 99, 101, 5, 39, 20, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2,
	104, 108, 7, 48, 2, 2, 105, 107, 5, 39, 20, 2, 106, 105, 3, 2, 2, 2, 107,
	110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 118,
	3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 113, 7, 48, 2, 2, 112, 114, 5, 39,
	20, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2,
	115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 100, 3, 2, 2, 2, 117,
	111, 3, 2, 2, 2, 118, 36, 3, 2, 2, 2, 119, 121, 9, 2, 2, 2, 120, 119, 3,
	2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 125, 8, 19, 2, 2, 125, 38, 3, 2, 2, 2, 126,
	127, 9, 3, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 9, 4, 2, 2, 129, 42, 3,
	2, 2, 2, 130, 135, 5, 41, 21, 2, 131, 134, 5, 41, 21, 2, 132, 134, 5, 39,
	20, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2,
	135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 44, 3, 2, 2, 2, 137, 135,
	3, 2, 2, 2, 13, 2, 84, 86, 97, 102, 108, 115, 117, 122, 133, 135, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'->'", "'<->'", "'<--'", "'<-chan'", "'{'", "','", "'}'", "'['",
	"']'", "'@'", "':'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "QUOTED_STRING",
	"INT", "FLOAT", "WS", "ID",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "QUOTED_STRING", "ESC", "INT", "FLOAT",
	"WS", "DIGIT", "LETTER", "ID",
}

type nmdLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewnmdLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *nmdLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnmdLexer(input antlr.CharStream) *nmdLexer {
	l := new(nmdLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "nmd.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// nmdLexer tokens.
const (
	nmdLexerT__0          = 1
	nmdLexerT__1          = 2
	nmdLexerT__2          = 3
	nmdLexerT__3          = 4
	nmdLexerT__4          = 5
	nmdLexerT__5          = 6
	nmdLexerT__6          = 7
	nmdLexerT__7          = 8
	nmdLexerT__8          = 9
	nmdLexerT__9          = 10
	nmdLexerT__10         = 11
	nmdLexerT__11         = 12
	nmdLexerT__12         = 13
	nmdLexerQUOTED_STRING = 14
	nmdLexerINT           = 15
	nmdLexerFLOAT         = 16
	nmdLexerWS            = 17
	nmdLexerID            = 18
)
