package main

import (
	"github.com/alecthomas/participle"
	"github.com/the4thamigo-uk/units"
)

type (
	AST struct {
		Package      string                   `"package" @Ident ";"`
		BaseUnits    []*BaseUnitDefinition    `"base" "(" { @@ ";" } ")"`
		DerivedUnits []*DerivedUnitDefinition `"derived" "(" { @@ ";" } ")"`
		Quantities   []*QuantityDefinition    `"quantity" "(" { @@ ";" } ")"`
		Operations   []*OperationDefinition   `"operation" "(" { @@ ";" } ")"`
	}

	BaseUnitDefinition struct {
		Name    string `@Ident "="`
		Literal string `@String`
	}

	DerivedUnitDefinition struct {
		Name       string            `@Ident "="`
		Expression *units.Expression `@@`
	}

	QuantityDefinition struct {
		Name           string           `@Ident`
		UnitExpression units.Expression `"(" @@? ")"`
	}

	OperationDefinition struct {
		Result   string `@Ident "="`
		Left     string `@Ident`
		Operator string `@("*" | "/" | "+" | "-")`
		Right    string `@Ident`
	}
)

func parse(s string) (*AST, error) {
	parser, err := participle.Build(&AST{})
	if err != nil {
		return nil, err
	}

	var ast AST
	err = parser.ParseString(s, &ast)
	if err != nil {
		return nil, err
	}
	return &ast, nil
}
