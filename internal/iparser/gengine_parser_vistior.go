package iparser

import parser "github.com/cookedsteak/gengine/internal/iantlr/alr2"

type GengineParserVisitor struct {
	parser.BasegengineVisitor
}

func NewGengineParserVisitor() *GengineParserVisitor {
	return &GengineParserVisitor{}
}
