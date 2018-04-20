package main

import "fmt"

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sayHi() (x, y string) {
	x = "hello"
	y = "vto"
	return
}

func feedMeRecursively(portion int, eaten int) int {
	eaten = portion + eaten

	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}

	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMeRecursively(portion, eaten)
}

func functionPass(fn func() string) string {
	return fn()
}

func main() {
	quantity, prize := getPrize()
	fmt.Printf("You won %v %v\n", quantity, prize)

	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("The result is %v\n", result)

	fmt.Println(sayHi())

	feedMeRecursively(1, 2)

	fn := func() string {
		return "function called"
	}

	fmt.Println(functionPass(fn))
}
