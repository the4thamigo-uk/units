package main

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`
package {{.Package}}

import (
	"github.com/the4thamigo-uk/units"
	"fmt"
	"math"
)

 // ensure we use math package to avoid unused imports error for some code generations
var _ = math.Abs(0)

// units
var (
	{{range $n, $u := .Units}}{{$n}} = units.MustParse(units.Parse("{{$u.String}}"))
	{{end}}
)

// quantities
type (
	{{range .Quantities}}{{.TypeName}} {{.BaseType}}
	{{end}}
)

// quantity units
var (
{{range .Quantities}}{{.UnitName}} = units.MustParse(units.Parse("{{.Unit.String}}"))
{{.BaseUnitName}} = units.MustParse(units.Parse("{{.BaseUnit.String}}"))
{{end}}
)

// quantity zero values
var (
{{range .Quantities}}{{.ZeroValueName}} = {{.Constructor}}(0)
{{end}}
)

// quantity slices
type (
{{range .Quantities}}{{.TypeName}}Slice []*{{.TypeName}}
{{.MapperName}} func(q *{{.TypeName}}) *{{.TypeName}}
{{.ReducerName}} func(acc *{{.TypeName}}, q *{{.TypeName}}) *{{.TypeName}}
{{end}}
)
// declared conversions between quantities
var (
{{range $cv := .Conversions}}{{$cv.ConverterName}} = units.MustConvert(units.NewConverter({{$cv.Left.BaseUnitName}},{{$cv.Right.BaseUnitName}}))
{{end}}
)

{{range $q := .Quantities}}
func {{.Constructor}}(val {{.BaseType}}) {{.TypeName}} {
	return {{.TypeName}}(val)
}

func {{.PtrConstructor}}(val {{.BaseType}}) *{{.TypeName}} {
	q := {{.Constructor}}(val)
	return &q
}

func {{.FromPtrConstructor}}(val *{{.BaseType}}) *{{.TypeName}} {
	if val == nil {
		return nil
	}
	return {{.PtrConstructor}}(*val)
}

func (q *{{.TypeName}}) Value() *{{.BaseType}} {
	if q == nil {
		return nil
	}
	v := {{.BaseType}}(*q)
	return &v
}

func (q *{{.TypeName}}) ValueOrDefault(dft {{.BaseType}}) {{.BaseType}} {
	if q == nil {
		return dft
	}
	return {{.BaseType}}(*q)
}

func (q *{{.TypeName}}) Convert(u units.Unit) (*{{.BaseType}}, error) {
	if q == nil {
		return nil, nil
	}
	u2 := q.BaseUnit().Divide(u)	
	if !u2.IsScalar() {
		return nil, fmt.Errorf("cannot convert '%s' to given units '%s'", q.Unit(), u)
	}
	v := {{.BaseType}}(u2.Scale() * float64(*q.Value()))
	return &v, nil
}

func (q *{{.TypeName}}) ConvertOrDefault(u units.Unit, dft {{.BaseType}}) (*{{.BaseType}}, error) {
	if q == nil {
		return &dft, nil
	}
	return q.Convert(u)
}

func (q *{{.TypeName}}) Unit() units.Unit {
	return {{.UnitName}}
}

func (q *{{.TypeName}}) BaseUnit() units.Unit {
	return {{.BaseUnitName}}
}

func (q *{{.TypeName}}) IsZero() bool {
	return q.Eq(&{{.ZeroValueName}})
}

func (q *{{.TypeName}}) Eq(q2 *{{.TypeName}}) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() == *q2.Value()
}

func (q *{{.TypeName}}) Gt(q2 *{{.TypeName}}) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() > *q2.Value()
}

func (q *{{.TypeName}}) GtEq(q2 *{{.TypeName}}) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() >= *q2.Value()
}

func (q *{{.TypeName}}) Lt(q2 *{{.TypeName}}) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() < *q2.Value()
}

func (q *{{.TypeName}}) LtEq(q2 *{{.TypeName}}) bool {
	if q == nil || q2 == nil {
		return false
	}
	return *q.Value() <= *q2.Value()
}

func (q *{{.TypeName}}) Between(q1, q2 *{{.TypeName}}) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.GtEq(q1) && q.LtEq(q2)
}

func (q *{{.TypeName}}) Inside(q1, q2 *{{.TypeName}}) bool {
	if q == nil || q1 == nil || q2 == nil {
		return false
	}
	return q.Gt(q1) && q.Lt(q2)
}

func (q *{{.TypeName}}) Negate() *{{.TypeName}} {
	if q == nil {
		return nil
	}
	return {{.PtrConstructor}}(-*q.Value())
}

func (q *{{.TypeName}}) Abs() *{{.TypeName}} {
	if q == nil {
		return nil
	}
	if q.GtEq(&{{.ZeroValueName}}) {
		return q
	}
	return q.Negate()
}

func (q *{{.TypeName}}) Min(q2 *{{.TypeName}}) *{{.TypeName}} {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Lt(q2) {
		return q
	}
	return q2
}

func (q *{{.TypeName}}) Max(q2 *{{.TypeName}}) *{{.TypeName}} {
	if q == nil || q2 == nil {
		return nil
	}
	if q.Gt(q2) {
		return q
	}
	return q2
}

func (qq {{.SliceName}}) Map(f {{.MapperName}}) {{.SliceName}} {
	out := make({{.SliceName}}, len(qq))
	for i, q := range qq {
		out[i] = f(q)
	}
	return out
}

func (qq {{.SliceName}}) Reduce(acc *{{.TypeName}}, f {{.ReducerName}}) *{{.TypeName}} {
	for _, q := range qq {
		acc = f(acc, q)
	}
	return acc
}

{{range $op := .Operations}}
func (q *{{$q.TypeName}}) {{$op.OperationSpec "q2"}} {
	if q == nil || q2 == nil {
		return nil
	}
	{{if or (eq $op.Operator "/") (eq $op.Operator "%")}}
	if q2.IsZero() {
		return nil
	}
	{{end}}
	return {{$op.Result.PtrConstructor}}({{$op.OperationImpl "q" "q2"}})
}
{{end}}

{{range $cv := .Conversions}}
func (q *{{$q.TypeName}}) To{{$cv.Right.ProperName}}() *{{$cv.Right.Name}} {
	if q == nil {
		return nil
	}
	return {{$cv.Right.FromPtrConstructor}}({{$cv.ConverterName}}.ConvertPtr(q.Value()))
}
{{end}}

{{end}}
`))

func generate(s *Model) (string, error) {
	var b bytes.Buffer
	err := tmpl.Execute(&b, s)
	if err != nil {
		return "", fmt.Errorf("failed to generate : %w", err)
	}
	src, err := format.Source(b.Bytes())
	if err != nil {
		return "", fmt.Errorf("failed to format generated code : %w", err)
	}
	return string(src), nil
}
