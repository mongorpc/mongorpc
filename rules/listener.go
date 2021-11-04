package rules

import (
	"github.com/mongorpc/mongorpc/rules/parser"
)

type MongoRPCRulesListener struct {
	*parser.BaseMongoRPCRulesListener
}

func NewMongoRPCRulesListener() *MongoRPCRulesListener {
	return new(MongoRPCRulesListener)
}
