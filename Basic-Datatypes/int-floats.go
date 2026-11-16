package datatypes

import (
	"fmt"
	"math"
)

var (
	normalInt       int   = 0
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
}

func IntFloatExps() {
	// adding int and a float
	fmt.Println("Adding int and a float ", float64(a)+i)

	// adding int and float but only adding non-decimal part
	fmt.Println("Adding int and a float but omit decimal part ", int64(i)+a)
}

func truncateFloatMethod1(num float64, precision float64) float64 {
	return math.Trunc(num/precision) * precision
}
