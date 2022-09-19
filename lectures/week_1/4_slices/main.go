package main

import "fmt"

func main() {
	var sl []int
	fmt.Println("Значение, длина, объем:", sl, len(sl), cap(sl))
	// Значение, длина, объем: [] 0 0

	// добавление элемента в слайс
	sl = append(sl, 100)
	fmt.Println("Значение, длина, объем:", sl, len(sl), cap(sl))
	// Значение, длина, объем: [100] 1 1

	sl = append(sl, 102)
	fmt.Println("Значение, длина, объем:", sl, len(sl), cap(sl))
	// Значение, длина, объем: [100 102] 2 2

	sl = append(sl, 103)
	sl = append(sl, 104)
	fmt.Println("Значение, длина, объем:", sl, len(sl), cap(sl))
	// Значение, длина, объем: [100 102 103 104] 4 4

	sl = append(sl, 105)
	fmt.Println("Значение, длина, объем:", sl, len(sl), cap(sl))
	// Значение, длина, объем: [100 102 103 104 105] 5 8

	// короткая инициализация
	sl2 := []int{10, 20, 30}
	fmt.Println(sl2)
	// [10 20 30]

	// добавить слайс в слайс
	sl = append(sl, sl2...)
	fmt.Println(sl)
	// [100 102 103 104 105 10 20 30]

	// создать слайс с нужной длиной сразу
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))
	// [0 0 0 0 0 0 0 0 0 0] 10 10

	// создать слайс с нужной длиной и размером
	slice4 := make([]int, 10, 15)
	fmt.Println(slice4, len(slice4), cap(slice4))
	// [0 0 0 0 0 0 0 0 0 0] 10 15

	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))
	// [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6] 16 30

	// внутри слайса - ссылка на массив, она копируется если просто присвоить
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)
	// [0 100500 0 0 0 0 0 0 0 0 1 2 3 4 5 6] [0 100500 0 0 0 0 0 0 0 0 1 2 3 4 5 6]

	//неправильная попытка скопировать слайс - он оппробует скопировать сколько влезет
	var slice6 []int
	copy(slice6, slice5)
	fmt.Println(slice6)
	// []

	//правильное копирование слайса
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)
	// [0 100500 0 0 0 0 0 0 0 0 1 2 3 4 5 6]

	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	// часть слайса [100500 0 0 0] [0 100500] [1 2 3 4 5 6]

	slice8 := append(slice7[:2], slice7[10:]...)
	fmt.Println("из кусков слайса", slice8)
	// из кусков слайса [0 100500 1 2 3 4 5 6]

	a := [...]int{5, 6, 7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("слайс из массива", sl8)
	// слайс из массива [5 8 7]

}
