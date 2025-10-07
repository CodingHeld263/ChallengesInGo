package mathematischeaufgaben

func CalcSumAndCountAllNumbersDivBy_2_Or_7(maxvalue int) (count int, sum int) {
	count = 0
	sum = 0
	for i := 1; i < maxvalue; i++ {
		var divby_2_and_7 bool = i%2 == 0 || i%7 == 0
		if divby_2_and_7 {
			count++
			sum += i
		}
	}
	return count, sum
}
