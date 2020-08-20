package units

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

type (
	Expression struct {
		Lhs      Term        `@@`
		Operator *string     `( @( "*" | "/")`
		Rhs      *Expression `@@)?`
	}

	Term struct {
		ScalarTerm *ScalarTerm `(@@ |`
		UnitTerm   *UnitTerm   `@@)?`
	}

	ScalarTerm struct {
		Scale float64 `(@Number)`
	}

	UnitTerm struct {
		Name     string `@Ident`
		Exponent *int   `("^" @Number)?`
	}
)

func Parse(s string) (Unit, error) {
	type Root struct {
		Expr *Expression `@@?`
	}
	var root Root
	err := ParseInto(&root, s)
	if err != nil {
		return nil, err
	}
	if root.Expr == nil {
		return Scalar(1), nil
	}
	return root.Expr.Unit()
}

func ParseInto(out interface{}, s string) error {
	lex := lexer.Must(ebnf.New(`
		Ident = (alpha | "_") { "_" | alpha | digit } .
		Number = [ "-" | "+" ] ("." | digit) { "." | digit } [ ("e"|"E") Number] .
		Punct = ";" | "+" | "-" | "*" | "/" | "^" | "%" | "\"" | "(" | ")" | "=" .
		Whitespace = " " | "\r" | "\t" | "\n" .
		alpha = "a"…"z" | "A"…"Z" .
		digit = "0"…"9" .
	`))

	parser, err := participle.Build(out,
		participle.Lexer(lex),
		participle.Elide("Whitespace"),
	)
	if err != nil {
		return err
	}

	return parser.ParseString(s, out)
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
	if t.ScalarTerm != nil {
		return t.ScalarTerm.Unit()
	}
	if t.UnitTerm != nil {
		return t.UnitTerm.Unit()
	}
	return Scalar(1)
}

func (st ScalarTerm) Unit() Unit {
	return Scalar(st.Scale)
}

func (t UnitTerm) Unit() Unit {
	exp := 1
	if t.Exponent != nil {
		exp = *t.Exponent
	}
	return NewUnit(t.Name, exp)
}
