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
			Unit() units.Unit
			{{range .Operations}}{{.FunctionSpec}}
			{{end}}}
	{{end}}
)

// quantity units
var (
{{range $q := .Quantities}}{{.UnitName}} = units.Must(units.Parse("{{$q.Unit.String}}"))
{{end}}
)

{{range $q := .Quantities}}
func New{{.InterfaceName}}(val float64) {{.InterfaceName}} {
	return {{.ValueName}}(val)
}

func (q {{.ValueName}}) Value() float64 {
	return float64(q)
}

func (q {{.ValueName}}) Unit() units.Unit {
	return {{.UnitName}}
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
