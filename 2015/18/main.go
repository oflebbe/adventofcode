package main

import "fmt"

type Stringer interface {
	String() string
}

type Array[T Stringer] struct {
	Values []T
	Num    int
}

func NewArray[T Stringer](num int) *Array[T] {
	values := make([]T, num*num)
	return &Array[T]{Values: values, Num: num}
}

func (a Array[T]) Get(i, j int) T {
	return a.Values[j*a.Num+i]
}

func (a *Array[T]) Set(i, j int,value  T ) {
	a.Values[j*a.Num+i] = value
}

func (a Array[T]) String() string {
	r := ""
	for j := 0; j < a.Num; j++ {
		for i := 0; i < a.Num; i++ {
			r += a.Values[j*a.Num+i].String()
		}
		r += "\n"
	}
	return r
}

type Bool bool

func (b Bool) String() string {
	return "#"
}

func parse(num int, st string) *Array[Bool] {
	result := NewArray[Bool](num)

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			result.Set(i, j, true)
		}
	}
	return result
}

func main() {
	fmt.Printf("%v\n", parse(4, "        "))
}
