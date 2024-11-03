# Types

## Numbers

* int8 - 8-bit signed integer
* int16 - 16-bit signed integer
* int32 - 32-bit signed integer
* int64 - 64-bit signed integer
* uint8 - 8-bit unsigned integer
* uint16 - 16-bit unsigned integer
* uint32 - 32-bit unsigned integer
* uint64 - 64-bit unsigned integer
* int - Both int and uint contain same size, either 32 or 64 bit.
* uint - Both int and uint contain same size, either 32 or 64 bit.
* rune - It is a synonym of int32 and also represent Unicode code points.
* byte - It is a synonym of uint8.
* uintptr - It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
* float32 - 32-bit IEEE 754 floating-point number
* float64 - 64-bit IEEE 754 floating-point number
* complex64 - Complex numbers which contain float32 as a real and imaginary component.
* complex128 - Complex numbers which contain float64 as a real and imaginary component.

## Booleans

* bool

## Strings

* string

--------------------------------------------------------

```go
package main

func main() {
	var variable1 string
	variable1 = "aaa"

	var variable2 string = "bbb"

	var variable3 = "ccc"

	variable4 := "ddd"
	
	variable5, variable6 := "eee", "fff"

	println(variable1) // aaa
	println(variable2) // bbb
	println(variable3) // ccc
	println(variable4) // ddd
	println(variable5) // eee
	println(variable6) // fff
}

```

