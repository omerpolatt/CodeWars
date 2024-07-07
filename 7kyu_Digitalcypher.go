package main

import (
	"strconv"
)

/*
String ifadenin harflerini key değerinin basmak değerleri ile kriptolojisini isteeyen fonksiyonumuz
( Our function that requests the cryptology of the letters of the string expression with the print values of the key value )
*/

func Encode(str string, key int) []int {

	result := []int{}

	strkey := strconv.Itoa(key) // basamak basamak key üzerinde gezinmek için
	keyLen := len(strkey)

	for i, char := range str {

		keyDigit := int(strkey[i%keyLen] - '0')
		encodedValue := int(char) - 'a' + 1 + keyDigit
		result = append(result, encodedValue)
	}

	return result
}
