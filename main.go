package main

func palindromeReverse(str string) bool {
	strReverse := reverse(str)
	return strReverse == str
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromeReverseEach(str string) bool {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func palindromeFromMiddle(str string) bool {
	runes := []rune(str)
	length := len(runes)
	middleI, middleJ := (length/2)-1, (length/2)+1
	if length%2 == 0 {
		middleJ -= 1
	}

	for i, j := middleI, middleJ; i >= 0; i, j = i-1, j+1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
