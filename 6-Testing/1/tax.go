package tax

import "time"

func CalculateTax(amount float64) float64 {
	time.Sleep(time.Millisecond )
	if amount == 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
