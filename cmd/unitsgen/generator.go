package main

import (
	"bytes"
	"go/format"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`
package {{.Package}}

import (
	"github.com/the4thamigo-uk/units"
	"fmt"
)

// units
var (
	{{range $n, $u := .Units}}{{$n}} = units.Must(units.Parse("{{$u.String}}"))
	{{end}}
)

// quantities
type (
	{{range .Quantities}}
		{{.TypeName}} {{.BaseType}}
		{{.InterfaceName}} interface {
			Value() {{.BaseType}}
			Convert(units.Unit) ({{.BaseType}}, error)
			Unit() units.Unit
			BaseUnit() units.Unit
			Eq(q2 {{.InterfaceName}}) bool
			Gt(q2 {{.InterfaceName}}) bool
			GtEq(q2 {{.InterfaceName}}) bool
			Lt(q2 {{.InterfaceName}}) bool
			LtEq(q2 {{.InterfaceName}}) bool
			Between(q1, q2 {{.InterfaceName}}) bool
			Inside(q1, q2 {{.InterfaceName}}) bool
			Abs() {{.InterfaceName}}
			Min(q2 {{.InterfaceName}}) {{.InterfaceName}}
			Max(q2 {{.InterfaceName}}) {{.InterfaceName}}
			{{range .Operations}}{{.FunctionSpec}}
			{{end}}}
	{{end}}
)

// quantity units
var (
{{range .Quantities}}{{.UnitName}} = units.Must(units.Parse("{{.Unit.String}}"))
{{.BaseUnitName}} = units.Must(units.Parse("{{.BaseUnit.String}}"))
{{end}}
)

// quantity zero values
var (
{{range .Quantities}}{{.ZeroValueName}} = New{{.InterfaceName}}(0)
{{end}}
)

{{range $q := .Quantities}}
func New{{.InterfaceName}}(val {{.BaseType}}) {{.InterfaceName}} {
	return {{.TypeName}}(val)
}

func (q {{.TypeName}}) Value() {{.BaseType}} {
	return {{.BaseType}}(q)
}

func (q {{.TypeName}}) Convert(u units.Unit) ({{.BaseType}}, error) {
	u2 := q.BaseUnit().Divide(u)	
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return {{.BaseType}}(u2.Scale() * float64(q.Value())), nil
}

func (q {{.TypeName}}) Unit() units.Unit {
	return {{.UnitName}}
}

func (q {{.TypeName}}) BaseUnit() units.Unit {
	return {{.BaseUnitName}}
}

func (q {{.TypeName}}) Eq(q2 {{.InterfaceName}}) bool {
	return q.Value() == q2.Value()
}

func (q {{.TypeName}}) Gt(q2 {{.InterfaceName}}) bool {
	return q.Value() > q2.Value()
}

func (q {{.TypeName}}) GtEq(q2 {{.InterfaceName}}) bool {
	return q.Value() >= q2.Value()
}

func (q {{.TypeName}}) Lt(q2 {{.InterfaceName}}) bool {
	return q.Value() < q2.Value()
}

func (q {{.TypeName}}) LtEq(q2 {{.InterfaceName}}) bool {
	return q.Value() <= q2.Value()
}

func (q {{.TypeName}}) Between(q1, q2 {{.InterfaceName}}) bool {
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q {{.TypeName}}) Inside(q1, q2 {{.InterfaceName}}) bool {
	return q.Gt(q1) && q.Lt(q2)
}

func (q {{.TypeName}}) Negate() {{.InterfaceName}} {
	return New{{.InterfaceName}}(-q.Value())
}

func (q {{.TypeName}}) Abs() {{.InterfaceName}} {
	if q.GtEq({{.ZeroValueName}}) {
		return q
	}
	return q.Negate()
}

func (q {{.TypeName}}) Min(q2 {{.InterfaceName}}) {{.InterfaceName}} {
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q {{.TypeName}}) Max(q2 {{.InterfaceName}}) {{.InterfaceName}} {
	if q.Gt(q2) {
		return q
	}
	return q2
}

{{range $op := .Operations}}
func (q {{$q.TypeName}}) {{$op.FunctionSpec}} {
	return New{{$op.Result.Name}}(q.Value() {{$op.Operator}} val.Value())
}
{{end}}
{{end}}
`))

func generate(s *Semantics) (string, error) {
	var b bytes.Buffer
	err := tmpl.Execute(&b, s)
	if err != nil {
		return "", err
	}
	src, err := format.Source(b.Bytes())
	if err != nil {
		return "", err
	}
	return string(src), nil
}
