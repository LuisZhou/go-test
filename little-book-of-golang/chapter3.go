// #Chapter 3 - Maps, Arrays and Slices
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// ## Array
	// In Go, arrays are fixed. Declaring an array requires that we specify the size, and once the size is specified, it cannot grow.
	// var scores [10]int
	// scores[0] = 339
	scores := [4]int{9001, 9333, 212, 33}
	fmt.Println(len(scores))
	for index, value := range scores {
		fmt.Println(index, value)
	}

	// this allocate a array space, but not slice
	scores_using_new := new([10]int)
	fmt.Println(len(scores_using_new)) // 10

	// ## Slices
	// A slice is a lightweight structure that wraps and represents a portion of an array.
	// scores_slice := []int{1, 4, 293, 4, 9}

	// We use make instead of new because there's more to creating a slice than just allocating the memory (which is what new does).
	// Specifically, we have to allocate the memory for the underlying array and also initialize the slice.
	// In the above, we initialize a slice with a length of 10 and a capacity of 10.
	// The length is the size of the slice, the capacity is the size of the underlying array.
	// Notice, slice is wrapper of the array.
	scores_slice_make := make([]int, 10)
	fmt.Println(len(scores_slice_make), cap(scores_slice_make)) // 10, 10

	scores_slice_make2 := make([]int, 0, 10)
	fmt.Println(len(scores_slice_make2), cap(scores_slice_make2)) // 0, 10
	//scores_slice_make2[7] = 9033                                // out of range.
	//fmt.Println(scores_slice_make2)

	// scores_slice_make2 = append(scores_slice_make2, 5) // append 5 to slice.
	// fmt.Println(scores_slice_make2)                    // prints [5], Println can print slice directly.

	// operation is [from:to], use the slice of the underly array excluding the to.
	// [from:] means to the end.
	// [:X] means start to X
	scores_slice_make2 = scores_slice_make2[0:8] // expand its length (not expand, but use different slice.)
	scores_slice_make2[7] = 9033
	fmt.Println(scores_slice_make2)                               //[0 0 0 0 0 0 0 9033]
	fmt.Println(len(scores_slice_make2), cap(scores_slice_make2)) // 8, 10

	// load and make is overload.
	// Go is a language that, to the frustration of some, makes use of features which aren't exposed for developers to use.

	// only way to expand the underlying array is use append.
	// It turns out that append is pretty special. If the underlying array is full,
	// it will create a new larger array and copy the values over (this is exactly how dynamic arrays work in PHP, Python, Ruby, JavaScript, ...).
	// This is why, in the example above that used append, we had to __re-assign__ the value returned by append to our scores variable:
	// append might have created a new value if the original had no more space.

	scores_slice_make3 := make([]int, 0, 5)
	c := cap(scores_slice_make3)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores_slice_make3 = append(scores_slice_make3, i)

		// if our capacity has changed,
		// Go had to grow our array to accommodate the new data
		if cap(scores_slice_make3) != c {
			c = cap(scores_slice_make3) // had to __re-assign__ the value returned by append to our scores variable
			fmt.Println(c)              // 5, 10, 20, 40 // 2x algorithm // no matter how opereration, just *2
		}
	}

	// because the length is 5, also
	// scores3 := make([]int, 5)
	// scores3 = append(scores3, 9332)
	// fmt.Println(scores3) // [0 0 0 0 0 9332]

	// names := []string{"leto", "jessica", "paul"} // not set then len and cap, but the init content infer it.
	// checks := make([]bool, 10) // len, cap the same. set the len and cap at the same time.
	// var names []string	// no len, cap, all is zero.
	// scores := make([]int, 0, 20) // set the len and cap seperate.

	// names := []string{"leto", "jessica", "paul"}
	// fmt.Println(len(names), cap(names)) // 3, 3

	// checks := make([]bool, 10)
	// fmt.Println(len(checks), cap(checks)) // 10, 10

	// var names []string
	// fmt.Println(len(names), cap(names)) // 10, 10

	// listen: array and slice is first class (base type), if you want its address, just use &

	// saiyans []*Saiyans, is a array of *Saiyans

	// However, in other languages, a slice is actually a new array with the values of the original copied over.
	// go just use the underlying array.

	scores_test_is_a_new_array := []int{1, 2, 3, 4, 5}
	slice := scores_test_is_a_new_array[2:4]
	//slice[0] = 999
	fmt.Println(slice) //[1 2 999 4 5]

	haystack := "the spice must flow"
	index_of_space := strings.Index(haystack[5:], " ") // search from 's' : 4
	fmt.Println("the index of ", index_of_space)       // this is index after 5

	scores_test_is_a_new_array = removeAtIndex(scores_test_is_a_new_array, 2)
	fmt.Println(scores_test_is_a_new_array)

	// copy
	scores_for_copy := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores_for_copy[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores_for_copy)

	worst := make([]int, 5)
	//copy(worst, scores_for_copy[:5])
	//copy(worst[2:4], scores_for_copy[:5])
	//copy(worst[2:4], scores_for_copy[1:5])
	//copy(worst, scores_for_copy[1:50])   // just adapt to the size of 'worst'.
	fmt.Println(worst)
}

// Just for unorder list.
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
