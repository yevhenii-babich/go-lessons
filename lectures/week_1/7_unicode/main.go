package main

import "fmt"

func main() {
	var symbol rune = 'a'
	autoSymbol := 'a' // int
	unicodeSymbol := '⌘'
	unicodeSymbolByNumber := '\u2318'
	println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)

	str1 := "Привіт світ!"
	fmt.Println("ua: ", str1, len(str1), len([]rune(str1)))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}

	str2 := "你好世界"
	fmt.Println("cn: ", str2, len(str2))
	for index, runeValue := range str2 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}
	println(str2[1])

	bin := []byte(str2)
	fmt.Println("binary cn: ", bin, len(bin))
	for idx, val := range bin {
		fmt.Printf("raw binary idx: %v, oct: %v, hex: %x\n", idx, val, val)
	}
}
