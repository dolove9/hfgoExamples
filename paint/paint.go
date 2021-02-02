package main

import "fmt"

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / 10
	//fmt.Printf("%.2f liters needed\n", area / 10.0)
	//return 0
}

func main() {
	fmt.Printf("%.2f liters needed\n", paintNeeded(3.0, 5.2))
	fmt.Printf("%.2f liters needed\n", paintNeeded(4.2, 3.0))
	fmt.Printf("%.2f liters needed\n", paintNeeded(5.2, 3.5))
	fmt.Printf("%.2f liters needed\n", paintNeeded(5.0, 3.3))
	//paintNeeded(4.2, 3.0)
	//paintNeeded(5.2, 3.5)
	//paintNeeded(5.0, 3.3)
}
