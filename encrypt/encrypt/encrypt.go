package encrypt

func Encrypt(last string) string {
	runes := []rune(last)
	n := len(runes)
	if runes[n-1] == 'z' {
		runes[n-1] = 'a'
		runes = append(runes, 'a')
	} else {
		runes[n-1]++
	}

	return string(runes)
}
