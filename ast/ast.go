package ast

import "github.com/KSSidhu/go-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program Node
/*  - Our root node
    - Program.Statements are a slice of AST nodes in our program
*/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// Let Statement
type LetStatement struct {
	Token token.Token // the token.Let Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // The token.IDENT Token
	Value string
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Return Statement
type ReturnStatement struct {
	Token       token.Token // token.Return Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
