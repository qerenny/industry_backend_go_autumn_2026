package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return s
	}

	shift = shift % n
	if shift < 0 {
		shift += n
	}
	result := string(runes[shift:]) + string(runes[:shift])

	return result
}
