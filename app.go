package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investimentAmout, years, expectedReturnRate float64

	outputText("Investment amount: ")
	// you need to pass the pointer
	fmt.Scan(&investimentAmout)

	outputText("Expected return value: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investimentAmout, years, expectedReturnRate)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
func outputText(text string) {
	fmt.Print(&text)
}

func calculateFutureValue(investimentAmout float64, years float64, expectedReturnRate float64) (float64, float64) {
	futureValue := investimentAmout * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, realFutureValue
}
