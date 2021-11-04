// Code generated from MongoRPCRules.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // MongoRPCRules

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMongoRPCRulesListener is a complete listener for a parse tree produced by MongoRPCRulesParser.
type BaseMongoRPCRulesListener struct{}

var _ MongoRPCRulesListener = &BaseMongoRPCRulesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMongoRPCRulesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMongoRPCRulesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMongoRPCRulesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMongoRPCRulesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterService is called when production service is entered.
func (s *BaseMongoRPCRulesListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseMongoRPCRulesListener) ExitService(ctx *ServiceContext) {}

// EnterNamespaceIdentifier is called when production namespaceIdentifier is entered.
func (s *BaseMongoRPCRulesListener) EnterNamespaceIdentifier(ctx *NamespaceIdentifierContext) {}

// ExitNamespaceIdentifier is called when production namespaceIdentifier is exited.
func (s *BaseMongoRPCRulesListener) ExitNamespaceIdentifier(ctx *NamespaceIdentifierContext) {}

// EnterNamespaceBlock is called when production namespaceBlock is entered.
func (s *BaseMongoRPCRulesListener) EnterNamespaceBlock(ctx *NamespaceBlockContext) {}

// ExitNamespaceBlock is called when production namespaceBlock is exited.
func (s *BaseMongoRPCRulesListener) ExitNamespaceBlock(ctx *NamespaceBlockContext) {}

// EnterMatchBlock is called when production matchBlock is entered.
func (s *BaseMongoRPCRulesListener) EnterMatchBlock(ctx *MatchBlockContext) {}

// ExitMatchBlock is called when production matchBlock is exited.
func (s *BaseMongoRPCRulesListener) ExitMatchBlock(ctx *MatchBlockContext) {}

// EnterAllowKey is called when production allowKey is entered.
func (s *BaseMongoRPCRulesListener) EnterAllowKey(ctx *AllowKeyContext) {}

// ExitAllowKey is called when production allowKey is exited.
func (s *BaseMongoRPCRulesListener) ExitAllowKey(ctx *AllowKeyContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseMongoRPCRulesListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseMongoRPCRulesListener) ExitComment(ctx *CommentContext) {}

// EnterMatcher is called when production matcher is entered.
func (s *BaseMongoRPCRulesListener) EnterMatcher(ctx *MatcherContext) {}

// ExitMatcher is called when production matcher is exited.
func (s *BaseMongoRPCRulesListener) ExitMatcher(ctx *MatcherContext) {}

// EnterAllow is called when production allow is entered.
func (s *BaseMongoRPCRulesListener) EnterAllow(ctx *AllowContext) {}

// ExitAllow is called when production allow is exited.
func (s *BaseMongoRPCRulesListener) ExitAllow(ctx *AllowContext) {}

// EnterGetPathVariable is called when production getPathVariable is entered.
func (s *BaseMongoRPCRulesListener) EnterGetPathVariable(ctx *GetPathVariableContext) {}

// ExitGetPathVariable is called when production getPathVariable is exited.
func (s *BaseMongoRPCRulesListener) ExitGetPathVariable(ctx *GetPathVariableContext) {}

// EnterPathVariable is called when production pathVariable is entered.
func (s *BaseMongoRPCRulesListener) EnterPathVariable(ctx *PathVariableContext) {}

// ExitPathVariable is called when production pathVariable is exited.
func (s *BaseMongoRPCRulesListener) ExitPathVariable(ctx *PathVariableContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseMongoRPCRulesListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseMongoRPCRulesListener) ExitArg(ctx *ArgContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseMongoRPCRulesListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseMongoRPCRulesListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterMemberArg is called when production memberArg is entered.
func (s *BaseMongoRPCRulesListener) EnterMemberArg(ctx *MemberArgContext) {}

// ExitMemberArg is called when production memberArg is exited.
func (s *BaseMongoRPCRulesListener) ExitMemberArg(ctx *MemberArgContext) {}

// EnterMemberArguments is called when production memberArguments is entered.
func (s *BaseMongoRPCRulesListener) EnterMemberArguments(ctx *MemberArgumentsContext) {}

// ExitMemberArguments is called when production memberArguments is exited.
func (s *BaseMongoRPCRulesListener) ExitMemberArguments(ctx *MemberArgumentsContext) {}

// EnterArgDeclaration is called when production argDeclaration is entered.
func (s *BaseMongoRPCRulesListener) EnterArgDeclaration(ctx *ArgDeclarationContext) {}

// ExitArgDeclaration is called when production argDeclaration is exited.
func (s *BaseMongoRPCRulesListener) ExitArgDeclaration(ctx *ArgDeclarationContext) {}

// EnterArgDeclarations is called when production argDeclarations is entered.
func (s *BaseMongoRPCRulesListener) EnterArgDeclarations(ctx *ArgDeclarationsContext) {}

// ExitArgDeclarations is called when production argDeclarations is exited.
func (s *BaseMongoRPCRulesListener) ExitArgDeclarations(ctx *ArgDeclarationsContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseMongoRPCRulesListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseMongoRPCRulesListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFieldReferenceWithIdentifier is called when production fieldReferenceWithIdentifier is entered.
func (s *BaseMongoRPCRulesListener) EnterFieldReferenceWithIdentifier(ctx *FieldReferenceWithIdentifierContext) {
}

// ExitFieldReferenceWithIdentifier is called when production fieldReferenceWithIdentifier is exited.
func (s *BaseMongoRPCRulesListener) ExitFieldReferenceWithIdentifier(ctx *FieldReferenceWithIdentifierContext) {
}

// EnterFieldReferenceWithMemberRef is called when production fieldReferenceWithMemberRef is entered.
func (s *BaseMongoRPCRulesListener) EnterFieldReferenceWithMemberRef(ctx *FieldReferenceWithMemberRefContext) {
}

// ExitFieldReferenceWithMemberRef is called when production fieldReferenceWithMemberRef is exited.
func (s *BaseMongoRPCRulesListener) ExitFieldReferenceWithMemberRef(ctx *FieldReferenceWithMemberRefContext) {
}

// EnterMemberFunction is called when production memberFunction is entered.
func (s *BaseMongoRPCRulesListener) EnterMemberFunction(ctx *MemberFunctionContext) {}

// ExitMemberFunction is called when production memberFunction is exited.
func (s *BaseMongoRPCRulesListener) ExitMemberFunction(ctx *MemberFunctionContext) {}

// EnterId is called when production id is entered.
func (s *BaseMongoRPCRulesListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseMongoRPCRulesListener) ExitId(ctx *IdContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterNumberExpression is called when production numberExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterNumberExpression(ctx *NumberExpressionContext) {}

// ExitNumberExpression is called when production numberExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitNumberExpression(ctx *NumberExpressionContext) {}

// EnterObjectReferenceExpression is called when production objectReferenceExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterObjectReferenceExpression(ctx *ObjectReferenceExpressionContext) {
}

// ExitObjectReferenceExpression is called when production objectReferenceExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitObjectReferenceExpression(ctx *ObjectReferenceExpressionContext) {
}

// EnterParenthesisExpression is called when production parenthesisExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production parenthesisExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterMemberReferenceExpression is called when production memberReferenceExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterMemberReferenceExpression(ctx *MemberReferenceExpressionContext) {
}

// ExitMemberReferenceExpression is called when production memberReferenceExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitMemberReferenceExpression(ctx *MemberReferenceExpressionContext) {
}

// EnterBooleanExpression is called when production booleanExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterBooleanExpression(ctx *BooleanExpressionContext) {}

// ExitBooleanExpression is called when production booleanExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitBooleanExpression(ctx *BooleanExpressionContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterLogicalExpression is called when production LogicalExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production LogicalExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterGetExpression is called when production getExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterGetExpression(ctx *GetExpressionContext) {}

// ExitGetExpression is called when production getExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitGetExpression(ctx *GetExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterStringExpression is called when production stringExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterStringExpression(ctx *StringExpressionContext) {}

// ExitStringExpression is called when production stringExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitStringExpression(ctx *StringExpressionContext) {}

// EnterNullExpression is called when production nullExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterNullExpression(ctx *NullExpressionContext) {}

// ExitNullExpression is called when production nullExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitNullExpression(ctx *NullExpressionContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterMemberFunctionExpression is called when production memberFunctionExpression is entered.
func (s *BaseMongoRPCRulesListener) EnterMemberFunctionExpression(ctx *MemberFunctionExpressionContext) {
}

// ExitMemberFunctionExpression is called when production memberFunctionExpression is exited.
func (s *BaseMongoRPCRulesListener) ExitMemberFunctionExpression(ctx *MemberFunctionExpressionContext) {
}

// EnterObjectReference is called when production objectReference is entered.
func (s *BaseMongoRPCRulesListener) EnterObjectReference(ctx *ObjectReferenceContext) {}

// ExitObjectReference is called when production objectReference is exited.
func (s *BaseMongoRPCRulesListener) ExitObjectReference(ctx *ObjectReferenceContext) {}

// EnterGetPathExpressionVariable is called when production getPathExpressionVariable is entered.
func (s *BaseMongoRPCRulesListener) EnterGetPathExpressionVariable(ctx *GetPathExpressionVariableContext) {
}

// ExitGetPathExpressionVariable is called when production getPathExpressionVariable is exited.
func (s *BaseMongoRPCRulesListener) ExitGetPathExpressionVariable(ctx *GetPathExpressionVariableContext) {
}

// EnterGetPath is called when production getPath is entered.
func (s *BaseMongoRPCRulesListener) EnterGetPath(ctx *GetPathContext) {}

// ExitGetPath is called when production getPath is exited.
func (s *BaseMongoRPCRulesListener) ExitGetPath(ctx *GetPathContext) {}

// EnterRuleFunctionCall is called when production ruleFunctionCall is entered.
func (s *BaseMongoRPCRulesListener) EnterRuleFunctionCall(ctx *RuleFunctionCallContext) {}

// ExitRuleFunctionCall is called when production ruleFunctionCall is exited.
func (s *BaseMongoRPCRulesListener) ExitRuleFunctionCall(ctx *RuleFunctionCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseMongoRPCRulesListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseMongoRPCRulesListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMatchPath is called when production matchPath is entered.
func (s *BaseMongoRPCRulesListener) EnterMatchPath(ctx *MatchPathContext) {}

// ExitMatchPath is called when production matchPath is exited.
func (s *BaseMongoRPCRulesListener) ExitMatchPath(ctx *MatchPathContext) {}
