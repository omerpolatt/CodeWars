package main

/*
bonus miktarına göre karesel sayılarla price miktarının çarpımını bonus a göre karşılaştırıp kaçıncı seviyeye ulaşabileceğimizi veren program
( a program that compares the product of quadratic numbers and price amount according to the bonus amount and gives us how many levels we can reach )
*/

func Beeramid(bonus int, price float64) int {

	result := 0
	total := 0.0

	for i := 1; ; i++ {

		levelCost := float64(i*i) * price

		if total+levelCost <= float64(bonus) {
			total += levelCost
			result = i
		} else {
			break
		}
	}

	return result
}
