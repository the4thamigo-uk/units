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
		Operations []*Operation
	}

	Operation struct {
		Result   *Quantity
		Operator string
		Param    *Quantity
	}
)

func (q *Quantity) ValueName() string {
	return "_" + q.Name
}

func (q *Quantity) InterfaceName() string {
	return q.Name
}

func (q *Quantity) UnitName() string {
	return "_unit_" + q.Name
}
func (o *Operation) FunctionSpec() (string, error) {
	op, err := operatorName(o.Operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s(val %s) %s", op, o.Param.Name, o.Param.Name, o.Result.Name), nil
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
	}
	return "", fmt.Errorf("Operator '%s' not supported", op)
}

func evalQuantity(qd *QuantityDefinition) (*Quantity, error) {
	u, err := qd.UnitExpression.Unit()
	if err != nil {
		return nil, err
	}
	return &Quantity{
		Name: qd.Name,
		Unit: u,
	}, nil
}

func analyse(ast *AST) (*Semantics, error) {
	s := Semantics{
		Package:    ast.Package,
		Units:      UnitsMap{},
		Quantities: QuantitiesMap{},
	}
	for _, bu := range ast.BaseUnits {
		if _, ok := s.Units[bu.Name]; ok {
			return nil, fmt.Errorf("unit '%s' is defined more than once", bu.Name)
		}
		s.Units[bu.Name] = units.NewUnit(bu.Literal, 1)
	}

	for _, du := range ast.DerivedUnits {
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

	for _, qd := range ast.Quantities {
		q, err := evalQuantity(qd)
		if err != nil {
			return nil, fmt.Errorf("quantity '%s' is invalid : %w", qd.Name, err)
		}
		q.Unit = q.Unit.Subs(s.Units)

		err = q.Unit.Validate(s.Units)
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
