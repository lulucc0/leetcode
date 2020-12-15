package main

func lemonadeChange(bills []int) bool {

	d5, d10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			d5++
		} else if bill == 10 {
			d5--
			d10++
			if d5 < 0 {
				return false
			}
		} else {
			if d10 > 0 {
				d10--
				d5--
			} else {
				d5 -= 3
			}
			if d5 < 0 {
				return false
			}
		}
	}

	return true
}
