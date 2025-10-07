package mathematischeaufgaben

import (
	"fmt"
)

func IsEven(value int) {
	if value%2 == 0 {
		fmt.Println(value, " ist eine gerade Zahl!")
	} else {
		fmt.Println(value, "ist eine ungerade Zahl!")
	}
}
func IsOdd(value int) {
	if value%2 != 0 {
		fmt.Println(value, "ist eine ungerade Zahl!")
	} else {
		fmt.Println(value, " ist eine gerade Zahl!")
	}
}
