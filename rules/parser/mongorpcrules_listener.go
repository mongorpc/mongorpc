// Code generated from MongoRPCRules.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // MongoRPCRules

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MongoRPCRulesListener is a complete listener for a parse tree produced by MongoRPCRulesParser.
type MongoRPCRulesListener interface {
	antlr.ParseTreeListener

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterNamespaceIdentifier is called when entering the namespaceIdentifier production.
	EnterNamespaceIdentifier(c *NamespaceIdentifierContext)

	// EnterNamespaceBlock is called when entering the namespaceBlock production.
	EnterNamespaceBlock(c *NamespaceBlockContext)

	// EnterMatchBlock is called when entering the matchBlock production.
	EnterMatchBlock(c *MatchBlockContext)

	// EnterAllowKey is called when entering the allowKey production.
	EnterAllowKey(c *AllowKeyContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterMatcher is called when entering the matcher production.
	EnterMatcher(c *MatcherContext)

	// EnterAllow is called when entering the allow production.
	EnterAllow(c *AllowContext)

	// EnterGetPathVariable is called when entering the getPathVariable production.
	EnterGetPathVariable(c *GetPathVariableContext)

	// EnterPathVariable is called when entering the pathVariable production.
	EnterPathVariable(c *PathVariableContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterMemberArg is called when entering the memberArg production.
	EnterMemberArg(c *MemberArgContext)

	// EnterMemberArguments is called when entering the memberArguments production.
	EnterMemberArguments(c *MemberArgumentsContext)

	// EnterArgDeclaration is called when entering the argDeclaration production.
	EnterArgDeclaration(c *ArgDeclarationContext)

	// EnterArgDeclarations is called when entering the argDeclarations production.
	EnterArgDeclarations(c *ArgDeclarationsContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFieldReferenceWithIdentifier is called when entering the fieldReferenceWithIdentifier production.
	EnterFieldReferenceWithIdentifier(c *FieldReferenceWithIdentifierContext)

	// EnterFieldReferenceWithMemberRef is called when entering the fieldReferenceWithMemberRef production.
	EnterFieldReferenceWithMemberRef(c *FieldReferenceWithMemberRefContext)

	// EnterMemberFunction is called when entering the memberFunction production.
	EnterMemberFunction(c *MemberFunctionContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterNumberExpression is called when entering the numberExpression production.
	EnterNumberExpression(c *NumberExpressionContext)

	// EnterObjectReferenceExpression is called when entering the objectReferenceExpression production.
	EnterObjectReferenceExpression(c *ObjectReferenceExpressionContext)

	// EnterParenthesisExpression is called when entering the parenthesisExpression production.
	EnterParenthesisExpression(c *ParenthesisExpressionContext)

	// EnterArithmeticExpression is called when entering the arithmeticExpression production.
	EnterArithmeticExpression(c *ArithmeticExpressionContext)

	// EnterMemberReferenceExpression is called when entering the memberReferenceExpression production.
	EnterMemberReferenceExpression(c *MemberReferenceExpressionContext)

	// EnterBooleanExpression is called when entering the booleanExpression production.
	EnterBooleanExpression(c *BooleanExpressionContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterCompareExpression is called when entering the compareExpression production.
	EnterCompareExpression(c *CompareExpressionContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterLogicalExpression is called when entering the LogicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterGetExpression is called when entering the getExpression production.
	EnterGetExpression(c *GetExpressionContext)

	// EnterInExpression is called when entering the inExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterStringExpression is called when entering the stringExpression production.
	EnterStringExpression(c *StringExpressionContext)

	// EnterNullExpression is called when entering the nullExpression production.
	EnterNullExpression(c *NullExpressionContext)

	// EnterRangeExpression is called when entering the rangeExpression production.
	EnterRangeExpression(c *RangeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterMemberFunctionExpression is called when entering the memberFunctionExpression production.
	EnterMemberFunctionExpression(c *MemberFunctionExpressionContext)

	// EnterObjectReference is called when entering the objectReference production.
	EnterObjectReference(c *ObjectReferenceContext)

	// EnterGetPathExpressionVariable is called when entering the getPathExpressionVariable production.
	EnterGetPathExpressionVariable(c *GetPathExpressionVariableContext)

	// EnterGetPath is called when entering the getPath production.
	EnterGetPath(c *GetPathContext)

	// EnterRuleFunctionCall is called when entering the ruleFunctionCall production.
	EnterRuleFunctionCall(c *RuleFunctionCallContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterMatchPath is called when entering the matchPath production.
	EnterMatchPath(c *MatchPathContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitNamespaceIdentifier is called when exiting the namespaceIdentifier production.
	ExitNamespaceIdentifier(c *NamespaceIdentifierContext)

	// ExitNamespaceBlock is called when exiting the namespaceBlock production.
	ExitNamespaceBlock(c *NamespaceBlockContext)

	// ExitMatchBlock is called when exiting the matchBlock production.
	ExitMatchBlock(c *MatchBlockContext)

	// ExitAllowKey is called when exiting the allowKey production.
	ExitAllowKey(c *AllowKeyContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitMatcher is called when exiting the matcher production.
	ExitMatcher(c *MatcherContext)

	// ExitAllow is called when exiting the allow production.
	ExitAllow(c *AllowContext)

	// ExitGetPathVariable is called when exiting the getPathVariable production.
	ExitGetPathVariable(c *GetPathVariableContext)

	// ExitPathVariable is called when exiting the pathVariable production.
	ExitPathVariable(c *PathVariableContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitMemberArg is called when exiting the memberArg production.
	ExitMemberArg(c *MemberArgContext)

	// ExitMemberArguments is called when exiting the memberArguments production.
	ExitMemberArguments(c *MemberArgumentsContext)

	// ExitArgDeclaration is called when exiting the argDeclaration production.
	ExitArgDeclaration(c *ArgDeclarationContext)

	// ExitArgDeclarations is called when exiting the argDeclarations production.
	ExitArgDeclarations(c *ArgDeclarationsContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFieldReferenceWithIdentifier is called when exiting the fieldReferenceWithIdentifier production.
	ExitFieldReferenceWithIdentifier(c *FieldReferenceWithIdentifierContext)

	// ExitFieldReferenceWithMemberRef is called when exiting the fieldReferenceWithMemberRef production.
	ExitFieldReferenceWithMemberRef(c *FieldReferenceWithMemberRefContext)

	// ExitMemberFunction is called when exiting the memberFunction production.
	ExitMemberFunction(c *MemberFunctionContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitNumberExpression is called when exiting the numberExpression production.
	ExitNumberExpression(c *NumberExpressionContext)

	// ExitObjectReferenceExpression is called when exiting the objectReferenceExpression production.
	ExitObjectReferenceExpression(c *ObjectReferenceExpressionContext)

	// ExitParenthesisExpression is called when exiting the parenthesisExpression production.
	ExitParenthesisExpression(c *ParenthesisExpressionContext)

	// ExitArithmeticExpression is called when exiting the arithmeticExpression production.
	ExitArithmeticExpression(c *ArithmeticExpressionContext)

	// ExitMemberReferenceExpression is called when exiting the memberReferenceExpression production.
	ExitMemberReferenceExpression(c *MemberReferenceExpressionContext)

	// ExitBooleanExpression is called when exiting the booleanExpression production.
	ExitBooleanExpression(c *BooleanExpressionContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitCompareExpression is called when exiting the compareExpression production.
	ExitCompareExpression(c *CompareExpressionContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitLogicalExpression is called when exiting the LogicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitGetExpression is called when exiting the getExpression production.
	ExitGetExpression(c *GetExpressionContext)

	// ExitInExpression is called when exiting the inExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitStringExpression is called when exiting the stringExpression production.
	ExitStringExpression(c *StringExpressionContext)

	// ExitNullExpression is called when exiting the nullExpression production.
	ExitNullExpression(c *NullExpressionContext)

	// ExitRangeExpression is called when exiting the rangeExpression production.
	ExitRangeExpression(c *RangeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitMemberFunctionExpression is called when exiting the memberFunctionExpression production.
	ExitMemberFunctionExpression(c *MemberFunctionExpressionContext)

	// ExitObjectReference is called when exiting the objectReference production.
	ExitObjectReference(c *ObjectReferenceContext)

	// ExitGetPathExpressionVariable is called when exiting the getPathExpressionVariable production.
	ExitGetPathExpressionVariable(c *GetPathExpressionVariableContext)

	// ExitGetPath is called when exiting the getPath production.
	ExitGetPath(c *GetPathContext)

	// ExitRuleFunctionCall is called when exiting the ruleFunctionCall production.
	ExitRuleFunctionCall(c *RuleFunctionCallContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitMatchPath is called when exiting the matchPath production.
	ExitMatchPath(c *MatchPathContext)
}
