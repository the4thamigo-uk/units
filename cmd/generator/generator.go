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
	{{range .Units}}{{if .Text}}{{.Name}} = units.NewUnit("{{.Text}}");
	{{else}}{{.Name}} = {{.Expression.AsCode}};
	{{end}}{{end}}
)

// quantities
type (
	{{range .Quantities}}
		{{.PrivateName}} float64
		{{.PublicName}} interface {
			Value() float64
			Unit() units.Unit
			{{range .OperationDefinitions}}{{.FunctionSpec}}
			{{end}}
		}
	{{end}}
)

{{range $q := .Quantities}}
func New{{.PublicName}}(val float64) {{.PublicName}} {
	return {{.PrivateName}}(val)
}

func (q {{.PrivateName}}) Value() float64 {
	return float64(q)
}

func (q {{.PrivateName}}) Unit() units.Unit {
	return {{.Unit}}
}

{{range $op := .OperationDefinitions}}
func (q {{$q.PrivateName}}) {{$op.FunctionSpec}} {
	return New{{$op.Result}}(q.Value() {{$op.Expression.Operator}} val.Value())
}
{{end}}
{{end}}
`))

func generate(f *File) (string, error) {
	var b bytes.Buffer
	err := tmpl.Execute(&b, f)
	if err != nil {
		return "", err
	}
	src, err := format.Source(b.Bytes())
	if err != nil {
		return "", err
	}
	return string(src), nil
}
