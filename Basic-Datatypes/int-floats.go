package datatypes

import (
	"encoding/json"
	"fmt"
	"math"
)

var (
	normalInt       int   = 0 // int can either be 32 or 64 bit based on sys architecture. Its better to use int32 or int64 in standard projects.
	eightBitInt     int8  = 21
	sixteenBitInt   int16 = 4567
	thirtyTwoBitInt int32 = 767986786
	sixtyFourBitInt int64 = 1234345466567
	a               int64 = 25
	b               int64 = 34
)

var (
	i float64 = 23.12
	j float64 = 28.897
)

func NumExps() {
	OnlyIntExps()
	OnlyFloatExps()
	IntFloatExps()
}

func OnlyIntExps() {
	// different int types

	fmt.Printf("Normal Int %v of type %T \n", normalInt, normalInt)
	fmt.Printf("8 Int %v of type %T \n", eightBitInt, eightBitInt)
	fmt.Printf("16 Int %v of type %T \n", sixteenBitInt, sixteenBitInt)
	fmt.Printf("32 Int %v of type %T \n", thirtyTwoBitInt, thirtyTwoBitInt)
	fmt.Printf("64 Int %v of type %T \n", sixtyFourBitInt, sixtyFourBitInt)

	fmt.Printf("Implicit Int type is %T \n", 1232)

	// adding two ints
	fmt.Println("Addig two ints ", a+b)

	maxIntVal := math.MaxInt64
	//overflow does not  panic and returns negative max - offset
	fmt.Println("overflow of ints does not cause panic ", maxIntVal+3, math.MinInt64) //returns -maxValue + (offset-1)(2 in this case)
}

func OnlyFloatExps() {
	// adding two floats
	fmt.Println("Adding two floats ", i+j)

	// adding two floats but restricting to 2 decimals
	fmt.Printf("Adding two floats but setting precision to 2: %.2f \n", truncateFloatMethod1(i, 0.01)+truncateFloatMethod1(j, 0.01))

	// rouding off float64 to nearest int
	fmt.Println("rouding off float ", math.Round(j))
	fmt.Println("ceiling off float ", math.Ceil(j))
	fmt.Println("flooring off float ", math.Floor(j))

	fmt.Println("rounding off negative decimals ", math.Round(-25.78))
	fmt.Println("ceiling off float ", math.Ceil(-25.78))
	fmt.Println("flooring off float ", math.Floor(-25.78))

	var a, b, c float64 = 0.1, 0.2, 0.3
	fmt.Println(a+b == c) // This returns false as 0.1,0.2,0.3 does not have finite representation in binary

	var first, second, third float64
	first = 10.0
	second = 11.0
	third = 21.0
	fmt.Println(first+second == third) // This returns true as 10.0, 11.0, 21.0 have finite representation in binary

	fmt.Println("Best way to check for equality in floats is to check diff tolerance ", math.Abs(c-b) <= 0.11)
}

func IntFloatExps() {
	// adding int and a float
	fmt.Println("Adding int and a float ", float64(a)+i)

	// adding int and float but only adding non-decimal part
	fmt.Println("Adding int and a float but omit decimal part ", int64(i)+a)

	// Any number defaults to float64 during JSON decode
	type Json struct {
		Num any `json:"anyNum"`
	}
	var jsonStruct Json
	jsonString :=
		`
		{
			"anyNum": 23
		}
	`
	json.Unmarshal([]byte(jsonString), &jsonStruct)
	fmt.Printf("Value of variable is %v and its type is %T\n", jsonStruct.Num, jsonStruct.Num) //By declaring 'Num' type as any, json decodes it as float64 instead of any int types.
}

func truncateFloatMethod1(num float64, precision float64) float64 {
	return math.Trunc(num/precision) * precision
}
