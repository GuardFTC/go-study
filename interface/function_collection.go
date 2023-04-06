// Package _interface @Author:冯铁城 [17615007230@163.com] 2023-04-02 10:17:33
package _interface

import (
	"fmt"
)

type appender interface {
	append(int)
}

type lener interface {
	len() int
}

type List []int

func (l *List) append(val int) {
	*l = append(*l, val)
}

func (l List) len() int {
	return len(l)
}

func countInto(a appender, start, end int) {
	for i := start; i <= end; i++ {
		a.append(i)
	}
}

func longEnough(l lener) bool {
	return l.len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type appender in argument to CountInto:
	//       List does not implement appender (Append method has pointer receiver)
	//countInto(lst, 1, 10)
	if longEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	countInto(plst, 1, 10) //VALID:Identical receiver type
	if longEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}
