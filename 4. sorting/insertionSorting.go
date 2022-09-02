package main

import "fmt"

func main() {

	sequence := []int{31, 41, 59, 26, 41, 58}
	sort(sequence)
	fmt.Println(sequence)
}

func sort(elements []int) {
	var n = len(elements)
	var i int

	for i = 1; i < n; i++ {
		var j int
		j = i

		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}

//input = a sequence of n numbers { a1,a2,.....,an}
//output = a reordering of the input sequence such that a1 < a2,<....<an
