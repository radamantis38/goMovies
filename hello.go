package main

import "errors"

func Sum(x int, y int) (int, error) {
	return 0, errors.New("Error")
}

func main() {
	Sum(5, 5)
}