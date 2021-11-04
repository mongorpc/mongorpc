package rules

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mongorpc/mongorpc/rules/parser"
)

type MongoRPCRules struct {
	FilePath string
}

func (r *MongoRPCRules) GetFilePath() string {
	return r.FilePath
}

func (r *MongoRPCRules) LoadRules() error {

	input, err := antlr.NewFileStream(r.FilePath)
	if err != nil {
		return err
	}
	lexer := parser.NewMongoRPCRulesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewMongoRPCRulesParser(stream)
	p.BuildParseTrees = true
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	service := p.Service()
	// ParseTreeWalker.DEFAULT.walk(this, service)
	antlr.ParseTreeWalkerDefault.Walk(NewMongoRPCRulesListener(), service)

	return nil
}
