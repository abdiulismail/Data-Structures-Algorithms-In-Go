package main

import "fmt"

func main() {
	square, cube, _ := powerSeries(4)
	fmt.Println("Square ", square, "Cube", cube)
}

//gets the power series of integer a and returns tuple of square of a cube of a
func powerSeries(a int) (int, int, error) {
	return a * a, a * a * a, nil
}
