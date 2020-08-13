package main

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/the4thamigo-uk/units"
	"unicode"
)

type (
	File struct {
		Package    string               `"package" @Ident ";"`
		Units      UnitDefinitions      `"unit" "(" { @@ } ")"`
		Quantities QuantityDefinitions  `"quantity" "(" { @@ } ")"`
		Operations OperationDefinitions `"operation" "(" { @@ } ")"`
		units      map[string]*UnitDefinition
		quantities map[string]*QuantityDefinition
	}

	UnitDefinitions      []*UnitDefinition
	QuantityDefinitions  []*QuantityDefinition
	OperationDefinitions []*OperationDefinition

	UnitDefinition struct {
		Name       string         `@Ident "=" `
		Text       *string        `(@String`
		Expression UnitExpression `| @@) ";"`
		Unit       units.Unit
	}

	UnitExpression struct {
		Left     string `@Ident`
		Operator string `@("*" | "/")`
		Right    string `@Ident`
	}

	QuantityDefinition struct {
		Name                 string `@Ident`
		Unit                 string `"(" @Ident ")" ";"`
		OperationDefinitions []*OperationDefinition
		UnitDefinition       *UnitDefinition
	}

	OperationDefinition struct {
		Result         string              `@Ident "="`
		Expression     OperationExpression `@@ ";"`
		UnitDefinition *UnitDefinition
	}

	OperationExpression struct {
		Left     string `@Ident`
		Operator string `@("*" | "/" | "+" | "-")`
		Right    string `@Ident`
	}
)

func makePrivate(name string) string {
	rs := []rune(name)
	rs[0] = unicode.ToLower(rs[0])
	return "_" + string(rs)
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

func (f *File) findUnit(name string) (*UnitDefinition, error) {
	ud, ok := f.units[name]
	if !ok || ud == nil {
		return nil, fmt.Errorf("failed to find unit '%s'", name)
	}
	return ud, nil
}

func (f *File) findQuantity(name string) (*QuantityDefinition, error) {
	qd, ok := f.quantities[name]
	if !ok || qd == nil {
		return nil, fmt.Errorf("failed to find quantity '%s'", name)
	}
	return qd, nil
}

func (f *File) init() error {
	f.initMaps()
	err := f.evalUnits()
	if err != nil {
		return err
	}
	err = f.evalQuantityUnits()
	if err != nil {
		return err
	}
	err = f.evalOperationUnits()
	if err != nil {
		return err
	}
	return nil
}

func (f *File) initMaps() {
	f.units = map[string]*UnitDefinition{}
	for _, ud := range f.Units {
		f.units[ud.Name] = ud
	}

	f.quantities = map[string]*QuantityDefinition{}
	for _, qd := range f.Quantities {
		f.quantities[qd.Name] = qd
		for _, od := range f.Operations {
			if od.Expression.Left == qd.Name {
				qd.OperationDefinitions = append(qd.OperationDefinitions, od)
			}
		}
	}
}

func (f *File) evalUnits() error {
	for _, ud := range f.Units {
		if ud.Text != nil {
			ud.Unit = units.NewUnit(*ud.Text)
			continue
		}
		exp := ud.Expression
		leftUnit, err := f.findUnit(exp.Left)
		if err != nil {
			return fmt.Errorf("could not resolve unit expression for '%s' : %w", ud.Name, err)
		}
		rightUnit, err := f.findUnit(exp.Right)
		if err != nil {
			return fmt.Errorf("could not resolve unit expression for '%s' : %w", ud.Name, err)
		}
		switch exp.Operator {
		case "*":
			ud.Unit = leftUnit.Unit.Multiply(rightUnit.Unit)
		case "/":
			ud.Unit = leftUnit.Unit.Divide(rightUnit.Unit)
		default:
			return fmt.Errorf("Operator '%s' not supported", exp.Operator)
		}
	}
	return nil
}

func (f *File) evalQuantityUnits() error {
	for _, qd := range f.quantities {
		ud, err := f.findUnit(qd.Unit)
		if err != nil {
			return err
		}
		qd.UnitDefinition = ud
	}
	return nil
}

func (f *File) evalOperationUnits() error {
	// match up operations with the quantity definitions
	for _, op := range f.Operations {
		// verify that the units are correct for each operation
		leftQuantity, err := f.findQuantity(op.Expression.Left)
		if err != nil {
			return err
		}
		rightQuantity, err := f.findQuantity(op.Expression.Right)
		if err != nil {
			return err
		}
		resultQuantity, err := f.findQuantity(op.Result)
		if err != nil {
			return err
		}
		leftUnit := leftQuantity.UnitDefinition.Unit
		rightUnit := rightQuantity.UnitDefinition.Unit
		resultUnit := resultQuantity.UnitDefinition.Unit

		exp := op.Expression
		var u units.Unit
		switch exp.Operator {
		case "*":
			u = leftUnit.Multiply(rightUnit)
		case "/":
			u = leftUnit.Divide(rightUnit)
		case "+":
			fallthrough
		case "-":
			if leftUnit != rightUnit {
				return fmt.Errorf("operands must have same unit for operator '%s', left is '%s' and right is '%s'", exp.Operator, leftUnit, rightUnit)
			}
			u = leftUnit
		default:
			return fmt.Errorf("Operator '%s' not supported", exp.Operator)
		}
		if !u.Equal(resultUnit) {
			return fmt.Errorf("declared unit of operation does not match implied unit of operands, result is '%s', left is '%s' and right is '%s'", resultUnit, leftUnit, rightUnit)
		}
	}
	return nil
}

func (ue *UnitExpression) AsCode() (string, error) {
	opName, err := operatorName(ue.Operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s(%s)", ue.Left, opName, ue.Right), nil
}

func (qd *QuantityDefinition) PrivateName() string {
	return makePrivate(qd.Name)
}

func (qd *QuantityDefinition) PublicName() string {
	return qd.Name
}

func (od *OperationDefinition) FunctionSpec() (string, error) {
	opName, err := operatorName(od.Expression.Operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s(val %s) %s", opName, od.Expression.Right, od.Expression.Right, od.Result), nil
}

func (oe *OperationExpression) OperatorName() (string, error) {
	switch oe.Operator {
	case "+":
		return "Add", nil
	case "-":
		return "Subtract", nil
	case "*":
		return "Multiply", nil
	case "/":
		return "Divide", nil
	}
	return "", fmt.Errorf("Operator '%s' not supported", oe.Operator)
}

func parse(cfg string) (*File, error) {
	parser, err := participle.Build(&File{})
	if err != nil {
		return nil, err
	}

	var f File
	err = parser.ParseString(cfg, &f)
	if err != nil {
		return nil, err
	}

	err = f.init()
	if err != nil {
		return nil, err
	}

	return &f, nil
}
