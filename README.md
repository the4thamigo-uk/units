# units

## Introduction

This is a PoC for a code generation tool that builds code that enforces correct units in calculations in statically type safe code.

## Getting Started

To build the example configuration file :

```
go build ./cmd/unitsgen && ./unitsgen -c ./example/example.units -o ./example/example.go
```

The tool works from a [configuration file](./example/example.units) that is divided into three parts :

The `unit` section defines the base units for the calculations and any units that can be derived from the base units by multiplication or division.
The names of units must follow the naming conventions of golang identifiers.
The following example defines two base units `m` and `s` and three derived units `km`, `h` and `kph` :

```
unit (
  m;
  s;
  mps = m / s;
  km = 1000 * m;
  h = s * 3600;
  kmph = km / h;
)
```

The `quantity` section defines the set of quantities that are required for the calculations along with their unit. The name of the quantity will
be the exact name of the type in the resulting golang code.
The following example defines threee quantities :

```
quantity (
  Length(km);
  Time(h);
  Speed(kmph);
)
```

Each `quantity` is represented by an `interface` in the golang code, and is provided with a constructor function and a few default methods
like this :

```
type Length  interface {
  Value() float64
  Unit() units.Unit
  ...
}

func NewLength(val float64) Length
```

The `operations` section defines the the operations that are allowed to perform on the quantities, to yield other quantities. Since this part
is not obvious to understand, here is an example of the operation to generate a `Speed`, from a `Length` and a `Time` :

```
Speed = Length / Time;
```

In terms of the golang code, this means that The `Length` quantity, will be generated with a method that looks something like this :

```
func (q _length) DivideTime(val Time) Speed {
	return NewSpeed(q.Value() / val.Value())
}
```

In terms of the calling code, you work like this :

```
	l := NewLength(100)
	t := NewTime(50)
	s := l.DivideTime(t)
  sKmph := s.Value()
  sMps, err := s.Convert(mps)
```

