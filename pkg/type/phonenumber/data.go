package phonenumber

func getNumbers(input string) string {
	var number string

	for _, t := range input {
		// 48 - 57 in ASCII this numbers 0-9
		if t >= 48 && t <= 57 {
			number += string(t)
		}
	}
	return number
}
