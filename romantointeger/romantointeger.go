package romantointeger

func RomanToInteger(s string) int {
	calculator := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	result := 0
	jump := false

	bytes := []byte(s)
	lastAvailable := len(bytes) - 1
	for i := range bytes {
		if jump {
			jump = false
			continue
		}
		if i != lastAvailable {
			if val, ok := calculator[string(bytes[i:i+2])]; ok {
				jump = true
				result += val
			}
		}
		if !jump {
			result += calculator[string(bytes[i])]
		}
	}
	return result
}
