package main

import "fmt"

func main() {
	num1 := 10
	num2 := 5
	result := countOperations(num1, num2)
	fmt.Println(result)
}
func countOperations(num1 int, num2 int) int {
	var count int = 0
	for num1 > 0 && num2 > 0 {

		if num1 >= num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		count++
	}
	//fmt.Println(count)
	return count

}
