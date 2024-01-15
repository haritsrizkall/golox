// generated code

package expression

import "github.com/haritsrizkall/golox/token"
  
type Expr struct{}

type Binary struct {
Left Expr
Token *token.Token
Right Expr
}
type Grouping struct {
Expression Expr
}
type Literal struct {
Value interface{}
}
type Unary struct {
Operator *token.Token
 Right Expr
}