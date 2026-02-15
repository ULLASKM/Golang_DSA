package patterns

import "fmt"

// Pattern-1: Rectangular Star Pattern
// Problem Statement: Given an integer N, print the following pattern

func RectangularStarPattern() {
	fmt.Println("Pattern-1: Rectangular Star Pattern")
	for range 5 {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Pattern-2: Right-Angled Triangle Pattern
// Problem Statement: Given an integer N, print the following pattern.

func TriangleStarPattern() {
	fmt.Println("Pattern-2: Right-Angled Triangle Pattern")
	for i := range 5 {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Pattern - 3: Right-Angled Number Pyramid
// Problem Statement: Given an integer N, print the following pattern :

func TriangleNumberPattern() {
	fmt.Println("Pattern - 3: Right-Angled Number Pyramid")
	for i := range 5 {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

// Pattern - 4: Right-Angled Number Pyramid - II
// Problem Statement: Given an integer N, print the following pattern :

func TriangleNumberPatternII() {
	fmt.Println("Pattern - 4: Right-Angled Number Pyramid - II")
	for i := range 5 {
		for j := 0; j <= i; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

// Pattern-5: Inverted Right Pyramid
// Problem Statement: Given an integer N, print the following pattern :

func InvertedTriangleStarPattern() {
	fmt.Println("Pattern-5: Inverted Right Pyramid")
	for i := 5; i > 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Pattern - 6: Inverted Numbered Right Pyramid
// Problem Statement: Given an integer N, print the following pattern :

func InvertedTriangleNumberPattern() {
	fmt.Println("Pattern - 6: Inverted Numbered Right Pyramid")
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

// Pattern - 7: Star Pyramid
// Problem Statement: Given an integer N, print the following pattern :

func PyramidStarPattern() {
	fmt.Println("Pattern - 7: Star Pyramid")
	for i := range 5 {
		for k := 5; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Pattern-8: Rectangular Star Hollow Pattern
// Problem Statement: Given an integer N, print the following pattern

func RectangularStarHollowPattern() {
	fmt.Println("Pattern-8: Rectangular Star Hollow Pattern")
	for i := range 5 {
		for j := 0; j < 5; j++ {
			if i == 0 || i == 4 {
				fmt.Print("*")
			} else if j == 0 || j == 4 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
