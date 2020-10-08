package main

import "fmt"

type Element struct {
	parent *Element
	size   int
	Data   interface{}
}

func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.parent = s
	s.size = 1
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e.parent != e {
		e = e.parent
	}
	return e
}

func FindWithRecursive(e *Element) *Element {
	if e.parent == e {
		return e
	} else {
		return FindWithRecursive(e.parent)
	}
}

func Union(e1, e2 *Element) {
	e1.parent = e2
}

func FastUnion(e1, e2 *Element) {
	e1SetName := Find(e1)
	e2SetName := Find(e2)
	if e1SetName == e2SetName {
		return
	}

	if e1SetName.size < e2SetName.size {
		e1SetName.parent = e2SetName
		e2SetName.size += e1SetName.size
	} else {
		e2SetName.parent = e1SetName
		e1SetName.size += e2SetName.size
	}
}

func main() {
	aSet := MakeSet("a")
	bSet := MakeSet("b")
	oneSet := MakeSet(1)
	twoSet := MakeSet(2)

	Union(aSet, bSet)
	Union(oneSet, twoSet)

	result := Find(aSet)
	fmt.Println(result.Data)

	result = Find(bSet)
	fmt.Println(result.Data)

	result = Find(oneSet)
	fmt.Println(result.Data)

	result = Find(twoSet)
	fmt.Println(result.Data)
}
