package main

import "fmt"

func main() {
OUTER:
	for {
		fmt.Println("Choose form the following operations:")
		fmt.Println("1. Addtion")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		var choice int
		fmt.Scan(&choice)
		if choice == 5 {
			break OUTER
		} else if choice < 1 || choice > 4 {
			fmt.Println("Invalid Choice")
			continue OUTER
		}
		fmt.Println("Enter 2 Numbers for performing operation")
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		switch choice {
		case 1:
			fmt.Printf("Addition Value %d\n", n1+n2)
			break
		case 2:
			fmt.Printf("Subtraction Value %d\n", n1-n2)
			break
		case 3:
			fmt.Printf("Multiplication Value %d\n", n1*n2)
			break
		case 4:
			fmt.Printf("Division Value %f\n", float32(n1)/float32(n2))
		}
		fmt.Println()
	}

}
