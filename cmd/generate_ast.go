package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Type struct {
	Name   string
	Fields string
}

var Types = []Type{
	{
		Name:   "Binary",
		Fields: "Left Expr,Token *token.Token,Right Expr",
	},
	{
		Name:   "Grouping",
		Fields: "Expression Expr",
	},
	{
		Name:   "Literal",
		Fields: "Value interface{}",
	},
	{
		Name:   "Unary",
		Fields: "Operator *token.Token, Right Expr",
	},
}

func main() {
	output := flag.String("o", "expression/generated_types.go", "Output file")

	flag.Parse()

	var b bytes.Buffer

	// generate header
	b.WriteString(`// generated code

package expression

import "github.com/haritsrizkall/golox/token"
  
type Expr struct{}
`)

	for _, typeExpr := range Types {
		b.WriteString("\n")
		fields := strings.Split(typeExpr.Fields, ",")

		fmt.Fprintf(&b, "type %s struct {\n", typeExpr.Name)

		b.WriteString(strings.Join(fields, "\n"))

		b.WriteString("\n}")
	}

	f, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(b.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println("Generating AST...")
}
