package main

import "fmt"

func main() {
LOOP:
	for no := 2; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Printf("Prime No : %d\n", no)
	}
	fmt.Println("Done!")
}
