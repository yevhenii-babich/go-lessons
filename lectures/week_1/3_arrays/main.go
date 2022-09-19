package main

import "fmt"

func main() {
	var arr1 [7]int // Масив заповнений значеннями за замовчуванням
	fmt.Printf("arr1 is a %T; len: %d; val: %v \n", arr1, len(arr1), arr1)
	// arr1 is a [7]int; len: 7; val: [0 0 0 0 0 0 0]

	// При декларації можна використовувати типізовану беззнакову константу
	const size uint = 3
	var arr2 [2 * size]bool
	fmt.Printf("arr2 is a %T; len: %d; val: %v \n", arr2, len(arr2), arr2)
	// arr2 is a [6]bool; len: 6; val: [false false false false false false]

	// Автоматичне визначення довжини під час заповнення значеннями
	arr3 := [...]int{7, 42, 11}
	fmt.Printf("arr3 is a %T; len: %d; val: %v \n", arr3, len(arr3), arr3)
	// arr3 is a [3]int; len: 3; val: [7 42 11]

	fmt.Println("Звернення за індексом:", arr3[2])
	// Звернення за індексом: 11

	arr3[1] = 12
	fmt.Println("Після зміни:", arr3)
	// Після зміни: [7 12 11]

	// не можна, перевірка під час компіляції
	// arr3[4] = 12
	// invalid array index 4 (out of bounds for 3-element array)

	var matrix [3][3]int
	matrix[1][1] = 1
	fmt.Println("Массив массивов:", matrix)
	// Масив масивів: [[0 0 0] [0 1 0] [0 0 0]]
	var xUnknown []int
	fmt.Printf("xUnknown: %v, len(%d), is nil: %v\n", xUnknown, len(xUnknown), xUnknown == nil)
	xUnknown = make([]int, 0)
	fmt.Printf("xUnknown: %v, len(%d), is nil: %v\n", xUnknown, len(xUnknown), xUnknown == nil)
}
