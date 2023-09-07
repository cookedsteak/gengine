package iparser

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type GengineErrorListener struct {
	antlr.ErrorListener
	GrammarErrors []string
}

func NewGengineErrorListener() *GengineErrorListener {
	return &GengineErrorListener{
		GrammarErrors: make([]string, 0),
	}
}

/*
*
syntax err check
*/
//func (el *GengineErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//	el.GrammarErrors = append(el.GrammarErrors, "line"+" "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
//}
//
//func (el *GengineErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//
//}
//func (el *GengineErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//
//}
//func (el *GengineErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//
//}
func (el *GengineErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	el.GrammarErrors = append(el.GrammarErrors, "line"+" "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func (el *GengineErrorListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (el *GengineErrorListener) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (el *GengineErrorListener) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ *antlr.ATNConfigSet) {
}
