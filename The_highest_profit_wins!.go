package main

/*
Verilen dizinin en küçük ve en büyük elemanını veren fonksiyonumuz
( Our function that returns the smallest and largest element of the given array )
*/
func MinMax(arr []int) [2]int {

	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return [2]int{min, max}
}
