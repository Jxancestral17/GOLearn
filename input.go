
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number between 1 - 10: ")
	numb, _ := reader.ReadString('\n')
	numbex, err := strconv.Atoi(numb[:len(numb)-1])
	if err != nil {
		fmt.Println(err)
	}
	if numbex < 1 || numbex > 10 {
		fmt.Println("The number is not between 1 and 10 ")
	}else{
    fmt.Println("Your number is ",numbex)
  }
}
