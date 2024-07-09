package main

/*
Verilen bir dizi içindeki sayıların en küçük ortak çarpanını bulmamızı sağlayan fonksiyonumuz
( Our function that allows us to find the least common factor of numbers in a given array )
*/
import (
	"math/big"
)

func min(nums []int64) int64 { // min sayıyı belirlediğimiz fonksiyon
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func LCM(nums ...int64) *big.Int {
	if len(nums) == 0 {
		return big.NewInt(1)
	}

	for _, num := range nums {
		if num == 0 {
			return big.NewInt(0)
		}
	}

	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	product := big.NewInt(1) // burada dizi içndeki üm sayıların çarpımını buluyoruz
	for _, num := range nums {
		product.Mul(product, big.NewInt(num))
	}

	for i := maxNum; i <= product.Int64(); i++ { // max sayı ile çarpımdaki sayıa ralığındaki tüm sayıları dizi içindeki tüm sayılar bölüyor mu onu buluyoruz
		// ve ilk bölen sayıy döndürüyoruz
		common := true
		for _, num := range nums {
			if i%num != 0 {
				common = false
				break
			}
		}
		if common {
			return big.NewInt(i)
		}
	}

	return product
}
