package main

import (
	"fmt"
	"github.com/the4thamigo-uk/units"
	"unicode"
)

type (
	Model struct {
		Package     string
		Units       UnitsMap
		Quantities  QuantitiesMap
		Conversions []*Conversion
	}

	UnitsMap      map[string]units.Unit
	QuantitiesMap map[string]*Quantity

	Quantity struct {
		Name        string
		Unit        units.Unit
		BaseUnit    units.Unit
		Operations  []*Operation
		Conversions []*Conversion
		BaseType    string
	}

	Operation struct {
		Result   *Quantity
		Operator string
		Left     *Quantity
		Right    *Quantity
	}

	Conversion struct {
		Left     *Quantity
		Operator string
		Right    *Quantity
	}
)

func (q *Quantity) ProperName() string {
	cc := []rune(q.Name)
	cc[0] = unicode.ToUpper(cc[0])
	return string(cc)
}

func (q *Quantity) TypeName() string {
	return q.Name
}

func (q *Quantity) SliceName() string {
	return q.Name + "Slice"
}

func (q *Quantity) MapperName() string {
	return q.Name + "Mapper"
}

func (q *Quantity) ReducerName() string {
	return q.Name + "Reducer"
}

func (q *Quantity) FilterName() string {
	return q.Name + "Filter"
}

func (q *Quantity) Constructor() string {
	return "New" + q.Name
}

func (q *Quantity) PtrConstructor() string {
	return q.Constructor() + "Ptr"
}

func (q *Quantity) FromPtrConstructor() string {
	return q.Constructor() + "FromPtr"
}

func (q *Quantity) UnitName() string {
	return "_unit_" + q.Name
}

func (q *Quantity) BaseUnitName() string {
	return "_base_unit_" + q.Name
}

func (q *Quantity) ZeroValueName() string {
	return "_zero_value_" + q.Name
}

func (q *Quantity) BaseTypeIsFloat() bool {
	return q.BaseType == "float32" || q.BaseType == "float64"
}

func (o *Operation) OperationSpec(prm string) (string, error) {
	op, err := operatorName(o.Operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s(%s *%s) *%s", op, o.Right.Name, prm, o.Right.Name, o.Result.Name), nil
}

func (o *Operation) OperationImpl(left, right string) (string, error) {
	op := o.Operator
	left = fmt.Sprintf("*%s.Value()", left)
	right = fmt.Sprintf("*%s.Value()", right)
	if op == "%" && (o.Left.BaseTypeIsFloat() || o.Right.BaseTypeIsFloat()) {
		return fmt.Sprintf("%s(math.Mod(float64(%s), float64(%s)))", o.Result.BaseType, left, right), nil
	}

	if o.Left.BaseType != o.Right.BaseType {
		right = fmt.Sprintf("%s(%s)", o.Left.BaseType, right)
	}

	impl := fmt.Sprintf("%s %s %s", left, o.Operator, right)

	if o.Result.BaseType != o.Left.BaseType {
		impl = fmt.Sprintf("%s(%s)", o.Result.BaseType, impl)
	}

	return impl, nil
}

func (c *Conversion) ConverterName() string {
	return fmt.Sprintf("_convert%sTo%s", c.Left.ProperName(), c.Right.ProperName())
}

func operatorName(op string) (string, error) {
	switch op {
	case "+":
		return "Add", nil
	case "-":
		return "Subtract", nil
	case "*":
		return "Multiply", nil
	case "/":
		return "Divide", nil
	case "%":
		return "Mod", nil
	}
	return "", fmt.Errorf("Operator '%s' not supported", op)
}

func evalQuantity(qd *QuantityDefinition) (*Quantity, error) {
	u, err := qd.UnitExpression.Unit()
	if err != nil {
		return nil, err
	}
	return &Quantity{
		Name:     qd.Name,
		Unit:     u,
		BaseType: qd.BaseType,
	}, nil
}

func buildModel(ast *AST) (*Model, error) {
	s := Model{
		Package:    ast.Package,
		Units:      UnitsMap{},
		Quantities: QuantitiesMap{},
	}
	for _, ud := range ast.Units {
		if ud.BaseUnit != nil {
			bu := ud.BaseUnit
			if _, ok := s.Units[bu.Name]; ok {
				return nil, fmt.Errorf("%s: unit '%s' is defined more than once", ud.Pos, bu.Name)
			}
			s.Units[bu.Name] = units.NewUnit(bu.Name, 1)
		} else if ud.DerivedUnit != nil {
			du := ud.DerivedUnit
			if _, ok := s.Units[du.Name]; ok {
				return nil, fmt.Errorf("%s: unit '%s' is defined more than once", ud.Pos, du.Name)
			}
			u, err := du.Expression.Unit()
			if err != nil {
				return nil, err
			}
			err = u.Validate(s.Units)
			if err != nil {
				return nil, fmt.Errorf("%s: failed to validate unit '%s' : %w", ud.Pos, du.Name, err)
			}
			s.Units[du.Name] = u.Subs(s.Units)
		}
	}

	for _, qd := range ast.Quantities {
		q, err := evalQuantity(qd)
		if err != nil {
			return nil, fmt.Errorf("%s: quantity '%s' is invalid : %w", qd.Pos, qd.Name, err)
		}
		q.BaseUnit = q.Unit.Subs(s.Units)

		err = q.BaseUnit.Validate(s.Units)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to validate unit for quantity '%s' : %w", qd.Pos, qd.Name, err)
		}
		s.Quantities[qd.Name] = q
	}

	for _, od := range ast.Operations {
		left, ok := s.Quantities[od.Left]
		if !ok {
			return nil, fmt.Errorf("%s: the left quantity '%s' is not defined for the operation", od.Pos, od.Left)
		}
		right, ok := s.Quantities[od.Right]
		if !ok {
			return nil, fmt.Errorf("%s: the right quantity '%s' is not defined for the operation", od.Pos, od.Right)
		}
		result, ok := s.Quantities[od.Result]
		if !ok {
			return nil, fmt.Errorf("%s: the result quantity '%s' is not defined for the operation", od.Pos, od.Result)
		}
		var bu units.Unit
		switch od.Operator {
		case "*":
			bu = left.BaseUnit.Multiply(right.BaseUnit)
		case "/":
			bu = left.BaseUnit.Divide(right.BaseUnit)
		case "%":
			fallthrough
		case "+":
			fallthrough
		case "-":
			if !left.Unit.Equal(right.Unit) {
				return nil, fmt.Errorf("%s: operations with '+' or '-' must have operands with the same unit, but %s=%s, and %s=%s", od.Pos, left.Name, left.Unit, right.Name, right.Unit)
			}
			bu = left.BaseUnit
		}
		if !bu.Equal(result.BaseUnit) {
			return nil, fmt.Errorf("%s: the unit generated by the operation '%s' does not match the unit '%s' of the resulting quantity", od.Pos, bu, result.BaseUnit)
		}

		left.Operations = append(left.Operations, &Operation{
			Result:   result,
			Operator: od.Operator,
			Left:     left,
			Right:    right,
		})
	}

	for _, cd := range ast.Conversions {
		left, ok := s.Quantities[cd.Left]
		if !ok {
			return nil, fmt.Errorf("%s: the left quantity '%s' is not defined for the conversion", cd.Pos, cd.Left)
		}
		right, ok := s.Quantities[cd.Right]
		if !ok {
			return nil, fmt.Errorf("%s: the right quantity '%s' is not defined for the conversion", cd.Pos, cd.Right)
		}
		c := &Conversion{
			Left:     left,
			Operator: cd.Operator,
			Right:    right,
		}
		left.Conversions = append(left.Conversions, c)
		s.Conversions = append(s.Conversions, c)

	}

	return &s, nil
}
