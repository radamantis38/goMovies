package main

import "errors"

func Sum(x int, y int) (int, error) {
	return 0, nil
}

func AmenorB (a int , b int) error{
	if a>=b {
		return errors.New("no es menor")
	}
	return nil
}

func main() {
	Sum(5, 5)
}