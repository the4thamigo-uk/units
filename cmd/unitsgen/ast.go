package main

import (
	"github.com/the4thamigo-uk/units"
)

type (
	AST struct {
		Package    string                 `"package" @Ident ";"`
		Units      []*UnitDefinition      `"unit" "(" { @@ ";" } ")"`
		Quantities []*QuantityDefinition  `"quantity" "(" { @@ ";" } ")"`
		Operations []*OperationDefinition `"operation" "(" { @@ ";" } ")"`
	}

	UnitDefinition struct {
		DerivedUnit *DerivedUnitDefinition `@@ |`
		BaseUnit    *BaseUnitDefinition    `@@`
	}

	BaseUnitDefinition struct {
		Name string `@Ident`
	}

	DerivedUnitDefinition struct {
		Name       string            `@Ident "="`
		Expression *units.Expression `@@`
	}

	QuantityDefinition struct {
		Name           string           `@Ident`
		UnitExpression units.Expression `"(" @@? ")"`
		BaseType       string           `@("int" | "int8" | "int16" | "int32" | "int64" | "uint" | "uint8" | "uint16" | "uint32" | "uint64" |"float32" | "float64")`
	}

	OperationDefinition struct {
		Result   string `@Ident "="`
		Left     string `@Ident`
		Operator string `@("*" | "/" | "+" | "-")`
		Right    string `@Ident`
	}
)

func parse(s string) (*AST, error) {
	var ast AST
	err := units.ParseInto(&ast, s)
	if err != nil {
		return nil, err
	}
	return &ast, nil
}
