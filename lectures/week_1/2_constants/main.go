package main

const (
	someInt        = 1
	typedInt int32 = 17
	fullName       = "Vasily"
)

const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota
	two
	_    // порожня змінна, перепустка iota
	four // = 4
)

const (
	_         = iota // пропускаємо перше значення
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// помилка! переповнення типу
	// ZB
)

func main() {
	pi := 3.14

	// тип константи може бути визначений під час компіляції
	println(pi + someInt)

	// константа може мати тип
	// println(pi + typedInt)
	// invalid operation: pi + typedInt (mismatched types float64 and int32)

	println(KB, MB, GB, TB, PB, EB)
}
