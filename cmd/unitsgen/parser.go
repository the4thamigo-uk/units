package main

import (
	"fmt"
	"github.com/the4thamigo-uk/units"
)

type (
	Semantics struct {
		Package    string
		Units      UnitsMap
		Quantities QuantitiesMap
	}

	UnitsMap      map[string]units.Unit
	QuantitiesMap map[string]*Quantity

	Quantity struct {
		Name       string
		Unit       units.Unit
		BaseUnit   units.Unit
		Operations []*Operation
		BaseType   string
	}

	Operation struct {
		Result   *Quantity
		Operator string
		Param    *Quantity
	}
)

func (q *Quantity) TypeName() string {
	return q.Name
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

func (o *Operation) FunctionSpec() (string, error) {
	op, err := operatorName(o.Operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s(q2 *%s) *%s", op, o.Param.Name, o.Param.Name, o.Result.Name), nil
}

func (o *Operation) FunctionImpl(left, right string) (string, error) {
	return fmt.Sprintf("%s %s %s", left, o.Operator, right), nil
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

func analyse(ast *AST) (*Semantics, error) {
	s := Semantics{
		Package:    ast.Package,
		Units:      UnitsMap{},
		Quantities: QuantitiesMap{},
	}
	for _, u := range ast.Units {
		if u.BaseUnit != nil {
			bu := u.BaseUnit
			if _, ok := s.Units[bu.Name]; ok {
				return nil, fmt.Errorf("unit '%s' is defined more than once", bu.Name)
			}
			s.Units[bu.Name] = units.NewUnit(bu.Name, 1)
		} else if u.DerivedUnit != nil {
			du := u.DerivedUnit
			if _, ok := s.Units[du.Name]; ok {
				return nil, fmt.Errorf("unit '%s' is defined more than once", du.Name)
			}
			u, err := du.Expression.Unit()
			if err != nil {
				return nil, err
			}
			err = u.Validate(s.Units)
			if err != nil {
				return nil, fmt.Errorf("failed to validate unit '%s' : %w", du.Name, err)
			}
			s.Units[du.Name] = u.Subs(s.Units)
		}
	}

	for _, qd := range ast.Quantities {
		q, err := evalQuantity(qd)
		if err != nil {
			return nil, fmt.Errorf("quantity '%s' is invalid : %w", qd.Name, err)
		}
		q.BaseUnit = q.Unit.Subs(s.Units)

		err = q.BaseUnit.Validate(s.Units)
		if err != nil {
			return nil, fmt.Errorf("failed to validate unit for quantity '%s' : %w", qd.Name, err)
		}
		s.Quantities[qd.Name] = q
	}

	for _, od := range ast.Operations {
		left, ok := s.Quantities[od.Left]
		if !ok {
			return nil, fmt.Errorf("the left quantity '%s' is not defined for the operation", od.Left)
		}
		right, ok := s.Quantities[od.Right]
		if !ok {
			return nil, fmt.Errorf("the right quantity '%s' is not defined for the operation", od.Right)
		}
		result, ok := s.Quantities[od.Result]
		if !ok {
			return nil, fmt.Errorf("the result quantity '%s' is not defined for the operation", od.Result)
		}
		var u units.Unit
		switch od.Operator {
		case "*":
			u = left.Unit.Multiply(right.Unit)
		case "/":
			u = left.Unit.Divide(right.Unit)
		case "+":
			fallthrough
		case "-":
			if !left.Unit.Equal(right.Unit) {
				return nil, fmt.Errorf("operations with '+' or '-' must have operands with the same unit, but %s=%s, and %s=%s", left.Name, left.Unit, right.Name, right.Unit)
			}
			u = left.Unit
		}
		if !u.Equal(result.Unit) {
			return nil, fmt.Errorf("the unit generated by the operation '%s' does not match the unit '%s' of the resulting quantity", u, result.Unit)
		}

		left.Operations = append(left.Operations, &Operation{
			Result:   result,
			Operator: od.Operator,
			Param:    right,
		})
	}
	return &s, nil
}
