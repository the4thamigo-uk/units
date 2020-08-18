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
		{{.ValueName}} float64
		{{.InterfaceName}} interface {
			Value() float64
			Convert(units.Unit) (float64, error)
			Unit() units.Unit
			BaseUnit() units.Unit
			{{range .Operations}}{{.FunctionSpec}}
			{{end}}}
	{{end}}
)

// quantity units
var (
{{range $q := .Quantities}}{{.UnitName}} = units.Must(units.Parse("{{$q.Unit.String}}"))
{{.BaseUnitName}} = units.Must(units.Parse("{{$q.BaseUnit.String}}"))
{{end}}
)

{{range $q := .Quantities}}
func New{{.InterfaceName}}(val float64) {{.InterfaceName}} {
	return {{.ValueName}}(val)
}

func (q {{.ValueName}}) Value() float64 {
	return float64(q)
}

func (q {{.ValueName}}) Convert(u units.Unit) (float64, error) {
	u2 := q.BaseUnit().Divide(u)	
	if !u2.IsScalar() {
		return 0, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	return u2.Scale() * float64(q), nil
}

func (q {{.ValueName}}) Unit() units.Unit {
	return {{.UnitName}}
}

func (q {{.ValueName}}) BaseUnit() units.Unit {
	return {{.BaseUnitName}}
}

{{range $op := .Operations}}
func (q {{$q.ValueName}}) {{$op.FunctionSpec}} {
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
