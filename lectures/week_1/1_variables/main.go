/*
багаторядковий коментар
змінні в golang
*/
package main

func main() {
	// целые числа
	var i int = 10 // platform-depended type, 32 or 64 bit
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1          // int8, int16, int32, int64
	var unsignedInt uint = 100500         // platform-depended type, 32 or 64 bit
	var unsignedBigInt uint64 = 1<<64 - 1 // uint8, unit16, uint32, unit64
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// numerics with floating-point
	var p float32 = 3.14 // float = float32, float64
	println("float: ", p)

	//Boolean
	var b = true
	println("bool variable", b)

	//String
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	// Бінарні дані
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	// так не можна
	// var singleQuote string = 'Hello world'
	// missing '
	// syntax error: unexpected це at end of statement

	/*
	   коротке оголошення
	*/
	meaningOfLive := 42
	println("Meaning of life is ", meaningOfLive)
	// працює тільки для нових змінних, world оголошено вище, тому помилка
	// world := "Світ"
	// no new variables on left side of :=

	/*
	   приведення типів
	*/
	println("float to int conversion ", int(p))
	println("int to string conversion ", string(48))

	// комплексні числа
	z := 2 + 3i
	println("complex number: ", z)

	/*
	   операції з рядками
	*/
	s1 := "Vasily"
	s2 := "Romanov"
	fullName := s1 + s2
	println("name length is: ", fullName, len(fullName))

	escaping := `Hello\r\n
	World`
	println("as-is escaping: ", escaping)

	/*
	   значення за замовчуванням
	*/
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	/*
		кілька змінних
	*/
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)
	println(m0, m2, m3)

}
