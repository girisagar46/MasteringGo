package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("denominator can't be zero")
	}
	return a / b, nil
}

func main() {
	for _, val := range [3]float64{2.0, 4.0, 0.0} {
		res, err := divide(90, val)
		if err != nil {
			fmt.Println("error", err)
		}
		fmt.Printf("%v/%v = %v\n", 90, val, res)
	}
}

/*
90/2 = 45
90/4 = 22.5
error denominator can't be zero
90/0 = 0
*/
