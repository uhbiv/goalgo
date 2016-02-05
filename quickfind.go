// Copyright 2015 github.com/uhbiv/goalgo
/*
	Algorithm: Quick Find
	Topics: Connected Components, Dynamic Connectivity, Connectivity Problem, Eager Algorithm

	Cost analysis
	init	|union	|find
	N	N	1

	Also See: http://algs4.cs.princeton.edu/15uf/
*/

package goalgo

// The main data structure for this implementation
var id []int

// QuickFindUF() will initialize N objects with as disjoint components
// N array accesses.
func QuickFindUF(N int) {
	count := N
	id := make([]int, count)

	for i := 0; i < len(id); i++ {
		id[i] = i
	}
}

// union() adds connection between two objects
// at most 2N+2 array accesses
// most expensive operation in the implementation
// for N union commands on N objects, its takes upto N^2 array accesses (quadratic time alert)
func Union(p, q int) {
	pid := id[p]
	qid := id[q]

	for i := range id {
		if id[i] == pid {
			id[i] = qid
		}
	}
}

// find() will return the component which an object belongs to
func Find(p int) int {
	return id[p]
}

// connected() checks whether two objects are in the same component
// Cost: 2 array accesses
func connected(p, q int) bool {
	return id[p] == id[q]
}

// count() returns the number of distinct components in the graph
// it simply iterates and adds to a map
// can also be used to return both the number of components
// and the length or elements of those components.
func count() int {
	d := make(map[int]int)

	for _, v := range id {
		d[v] += 1
	}
	return len(d)
}

// component() returns a list of the connected components and their objects
// return value type: map[int]int
// func component() map[int]int {
// 	d := make(map[int]int)

// 	for i, v := range id {
// 		d[v] += 1
// 	}
// 	return d
// }
