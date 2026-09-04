package ast

import "monke/token"

type Node interface {
	TokenLietral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLietral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLietral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLietral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (ls *Identifier) expressionNode()      {}
func (ls *Identifier) TokenLietral() string { return ls.Token.Literal }
