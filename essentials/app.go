package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investimentAmout, years, expectedReturnRate float64

	fmt.Print("Investment amount: ")
	// you need to pass the pointer
	fmt.Scan(&investimentAmout)

	fmt.Print("Expected return value: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investimentAmout * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
