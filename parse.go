package units

import (
	"fmt"
	"github.com/alecthomas/participle"
)

type (
	Expression struct {
		Lhs      Term        `@@`
		Operator *string     `( @( "*" | "/")`
		Rhs      *Expression `@@)?`
	}

	Term struct {
		Name     string `@Ident`
		Sign     string `("^" @("+" | "-")?`
		Exponent *int   `@Int)?`
	}
)

func Parse(s string) (Unit, error) {
	type Root struct {
		Expr *Expression `@@?`
	}

	parser, err := participle.Build(&Root{})
	if err != nil {
		return nil, err
	}

	var root Root
	err = parser.ParseString(s, &root)
	if err != nil {
		return nil, err
	}

	if root.Expr == nil {
		return Scalar(), nil
	}
	return root.Expr.Unit()
}

func Must(u Unit, err error) Unit {
	if err != nil {
		panic(err.Error())
	}
	return u
}

func (exp Expression) Unit() (Unit, error) {
	lhs := exp.Lhs.Unit()

	if exp.Operator == nil {
		return lhs, nil
	}

	rhs := exp.Rhs.Lhs.Unit()

	switch *exp.Operator {
	case "*":
		lhs = lhs.Multiply(rhs)
	case "/":
		lhs = lhs.Divide(rhs)
	default:
		return nil, fmt.Errorf("unit expression operator '%s' not supported", *exp.Operator)
	}
	rhs, err := Expression{
		Lhs:      Term{},
		Operator: exp.Rhs.Operator,
		Rhs:      exp.Rhs.Rhs,
	}.Unit()
	if err != nil {
		return nil, err
	}
	return lhs.Multiply(rhs), nil
}

func (t Term) Unit() Unit {
	exp := 1
	if t.Exponent != nil {
		exp = *t.Exponent
	}
	if t.Sign == "-" {
		exp = -exp
	}
	return NewUnit(t.Name, exp)
}
