// Return the Highest and Lowest number, start to array
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HL(in string) string {
	ins := strings.Fields(in)
	min := 0
	max := 0
	for _, value := range ins {
		var test, err = strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if min == 0 {
			min = test
		}
		if test < min {
			min = test
		}
		if test > max {
			max = test
		}

	}
	tot := strconv.Itoa(min) + " " + strconv.Itoa(max)
	return tot
}

func main() {
	fmt.Println(HL("1 2 3 4 5"))  // return "5 1"
	fmt.Println(HL("1 2 -3 4 5")) // return "5 -3"
	fmt.Println(HL("1 9 3 4 -5")) // return "9 -5"
}
