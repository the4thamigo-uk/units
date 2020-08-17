package units

import (
	"github.com/alecthomas/participle"
)

type (
	RootExpr struct {
		Expr *Expr `@@?`
	}

	Expr struct {
		Left     Term   `@@`
		Operator string `( @("*")`
		Right    *Expr  `@@)?`
	}

	Term struct {
		Unit     string `@Ident`
		Sign     string `( "^" @("-" | "+")?`
		Exponent *int   `@Int )?`
	}
)

func (re *RootExpr) toUnit() Unit {
	if re.Expr == nil {
		return Scalar()
	}

	return re.Expr.toUnit()
}

func (e *Expr) toUnit() Unit {
	u := e.Left.toUnit()
	switch e.Operator {
	case "":
		return u
	case "*":
		return u.Multiply(e.Right.toUnit())
	}
	panic("operator is not valid")
}

func (t *Term) toUnit() Unit {
	if t.Unit == "" {
		return Scalar()
	}
	exp := 1
	if t.Exponent != nil {
		exp = *t.Exponent
	}
	if t.Sign == "-" {
		exp *= -1
	}
	return unit{t.Unit: exp}
}

func Parse(s string) (Unit, error) {
	parser, err := participle.Build(&RootExpr{})
	if err != nil {
		return nil, err
	}

	var rootExpr RootExpr
	err = parser.ParseString(s, &rootExpr)
	if err != nil {
		return nil, err
	}

	return rootExpr.toUnit(), nil
}

func Must(u Unit, err error) Unit {
	if err != nil {
		panic(err.Error())
	}
	return u
}
