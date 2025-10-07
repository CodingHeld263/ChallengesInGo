package main

import (
	"fmt"

	mathematischeaufgaben "github.com/CodingHeld263/ChallengesInGo/Mathematische_Aufgaben"
)

func main() {
	fmt.Println(mathematischeaufgaben.Calc(6, 7))
	fmt.Println()
	fmt.Println(mathematischeaufgaben.CalcSumAndCountAllNumbersDivBy_2_Or_7(8))
	mathematischeaufgaben.IsEven(100)
	mathematischeaufgaben.IsOdd(100)
}
