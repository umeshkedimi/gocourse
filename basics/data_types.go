package basics

import "fmt"

func main() {
	fmt.Println("Hello" + " World")
	fmt.Println("9 x 10 =", 9*10)
	fmt.Println("180.18/2.0 =", 180.18/2.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(1 == 1)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
}
