package datatypes

import (
	"fmt"
	"math"
)

func IntFloatExps() {
	// adding two ints
	var a, b int64
	a = 25
	b = 34

	fmt.Println("Addig two ints ", a+b)

	// adding two floats
	var i, j float64
	i = 23.12
	j = 28.897

	fmt.Println("Adding two floats ", i+j)

	// adding two floats but restricting to 2 decimals
	fmt.Printf("Adding two floats but setting precision to 2: %.2f \n", TruncateFloatMethod1(i, 0.01)+TruncateFloatMethod1(j, 0.01))

	// adding int and a float
	fmt.Println("Adding int and a float ", float64(a)+i)

	// adding int and float but only adding non-decimal part
	fmt.Println("Adding int and a float but omit decimal part ", int64(i)+a)

	// rouding off float64 to nearest int
	fmt.Println("rouding off float ", math.Round(j))
	fmt.Println("ceiling off float ", math.Ceil(j))
	fmt.Println("flooring off float ", math.Floor(j))

	fmt.Println("rounding off negative decimals ", math.Round(-25.78))
	fmt.Println("ceiling off float ", math.Ceil(-25.78))
	fmt.Println("flooring off float ", math.Floor(-25.78))
}

func TruncateFloatMethod1(num float64, precision float64) float64 {
	return math.Trunc(num/precision) * precision
}
