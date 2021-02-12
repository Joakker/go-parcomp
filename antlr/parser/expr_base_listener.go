// Code generated from Expr.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseExprListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseExprListener) ExitFile(ctx *FileContext) {}

// EnterExprStmt is called when production ExprStmt is entered.
func (s *BaseExprListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production ExprStmt is exited.
func (s *BaseExprListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseExprListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseExprListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseExprListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseExprListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseExprListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseExprListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterNoOp is called when production NoOp is entered.
func (s *BaseExprListener) EnterNoOp(ctx *NoOpContext) {}

// ExitNoOp is called when production NoOp is exited.
func (s *BaseExprListener) ExitNoOp(ctx *NoOpContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseExprListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseExprListener) ExitPow(ctx *PowContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseExprListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseExprListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BaseExprListener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BaseExprListener) ExitExprlist(ctx *ExprlistContext) {}
