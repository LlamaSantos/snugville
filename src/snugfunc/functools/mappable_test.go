package functools

import (
	"fmt"
	"testing"
)

type Mappable interface {
	Map(func(interface{}) interface{}) interface{}
	Value() interface{}
}

type Box struct {
	value interface{}
}

func (b Box) Map(fn func(interface{}) interface{}) Box {
	return Box{value: fn(b.value)}
}

func (b Box) Value() interface{} {
	return b.value
}

func TestMappable(t *testing.T) {

	b := Box{value: 3}.Map(func(i interface{}) interface{} {
		return i.(int) + 2
	})

	fmt.Printf("Value: +%v\n", b.Value())
}
