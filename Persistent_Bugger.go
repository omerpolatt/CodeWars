package main

/*
Verilen sayının tek tek rakmalarını çarpıp bu çarpımın tek basamklı olana kadar devam etmsini sağlayıp kaçıncı adımda tek basamklı sayıya ulaştığımızı verdiren fonksiyonumuz 
*/

import "strconv"

func Persistence(n int) int {
	count := 0

	for n > 9 {
		çarpım := 1
		nStr := strconv.Itoa(n)

		for _, r := range nStr {
			digit, _ := strconv.Atoi(string(r))
			çarpım *= digit
		}

		n = çarpım
		count++
	}

	return count
}
