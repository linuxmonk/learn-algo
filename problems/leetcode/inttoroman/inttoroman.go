package inttoroman

var romanValues = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

var subtractions = map[int]string{
	4:   "IV",
	9:   "IX",
	40:  "XL",
	90:  "XC",
	400: "CD",
	900: "CM",
}

func IntToRoman(num int) string {
	v, ok := romanValues[num]
	if ok {
		return v
	}
	v, ok = subtractions[num]
	if ok {
		return v
	}
	var roman string
	var rep int
	var done bool
	for {
		if num == 0 || done {
			break
		}
		switch {
		case num >= 1000:
			rep = num / 1000
			num = num % 1000
			for i := 0; i < rep; i++ {
				roman += "M"
			}
		case num >= 900:
			num = num % 900
			roman += "CM"
		case num >= 500:
			num = num % 500
			roman += "D"
		case num >= 400:
			num = num % 400
			roman += "CD"
		case num >= 100:
			rep = num / 100
			num = num % 100
			for i := 0; i < rep; i++ {
				roman += "C"
			}
		case num >= 90:
			num = num % 90
			roman += "XC"
		case num >= 50:
			num = num % 50
			roman += "L"
		case num >= 40:
			num = num % 40
			roman += "XL"
		case num >= 10:
			rep = num / 10
			num = num % 10
			for i := 0; i < rep; i++ {
				roman += "X"
			}
		case num == 9:
			num = num % 9
			roman += "IX"
			done = true
			break
		case num >= 5:
			rep = num / 5
			num = num % 5
			for i := 0; i < rep; i++ {
				roman += "V"
			}
		case num == 4:
			done = true
			roman += "IV"
			break
		case num >= 1:
			for i := 0; i < num; i++ {
				roman += "I"
			}
			done = true
			break
		}
	}
	return roman
}
