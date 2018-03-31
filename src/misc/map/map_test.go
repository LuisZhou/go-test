package map_test

import (
	"sort"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	// where KeyType may be any type that is comparable (more on this later), and ValueType may be any type at all,
	// including another map!

	var m map[string]int

	// reference types: Map types are reference types, like pointers or slices, and so the value of m above is nil;
	// it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to
	// a nil map will cause a runtime panic; don't do that.

	// same as m = map[string]int{}.
	m = make(map[string]int)

	// ## Working with maps

	m["route"] = 66

	i := m["route"]

	j := m["root"]

	t.Log("value of i, j is ", i, j)

	n := len(m)

	t.Log("len of map is ", n)

	// the delete function doesn't return anything, and will do nothing if the specified key doesn't exist.
	delete(m, "route")

	i, ok := m["route"]
	t.Log(" A two-value assignment tests ", i, ok)

	m["1"] = 1
	m["2"] = 2

	for key, value := range m {
		t.Log("Key:", key, "Value:", value)
	}

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	for key, value := range commits {
		t.Log("Key:", key, "Value:", value)
	}

	// ## Exploiting zero values

	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] { // zero value for the boolean type is false
			t.Log("cycle detected")
			break
		}
		visited[n] = true
		t.Log(n.Value)
	}

	// another instance.

	// Another instance of helpful zero values is a map of slices. Appending to a nil slice just allocates a new slice,
	// so it's a one-liner to append a value to a map of slices; there's no need to check if the key exists. In the
	// following example, the slice people is populated with Person values. Each Person has a Name and a slice of Likes.
	// The example creates a map to associate each like with a slice of people that like it.

	type Person struct {
		Name  string
		Likes []string
	}
	var people []*Person

	peter := &Person{"peter", []string{"joe", "harden"}}
	jams := &Person{"jams", []string{"peter", "admen"}}
	green := &Person{"green", []string{"jams", "peter"}}

	people = []*Person{peter, jams, green}

	// double check: append.

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p) // Appending to a nil slice just allocates a new slice
		}
	}

	// Note that since both range and len treat a nil slice as a zero-length slice

	for _, p := range likes["peter"] {
		t.Log(p.Name, "likes peter.")
	}

	t.Log(len(likes["jams"]), "people like bacon.")

	// ## Key types

	// map keys may be of any type that is comparable:
	// boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types
	// absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys

	type Key struct {
		Path, Country string
	}
	hits := make(map[Key]int)
	hits[Key{"/", "vn"}]++ // if miss, just zero.
	hits[Key{"/ref/spec", "ch"}]++

	for k, v := range hits {
		t.Log(k, v)
	}

	// ## Concurrency
	// Why are map operations not defined to be atomic?
	// https://golang.org/doc/faq#atomic_maps
	// use sync.RWMutex

	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	counter.RLock()
	n1 := counter.m["some_key"]
	counter.RUnlock()
	t.Log("some_key:", n1)

	// ## Iteration order

	// When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the
	// same from one iteration to the next

	// If you require a stable iteration order you must maintain a separate data structure that specifies that order.

	var m1 map[int]string = map[int]string{
		1: "5",
		2: "6",
		3: "7",
		4: "8",
	}
	var keys []int
	for k := range m1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		t.Log("Key:", k, "Value:", m1[k])
	}
}
