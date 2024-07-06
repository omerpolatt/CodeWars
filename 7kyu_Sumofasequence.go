package main

/*
range döngü mantığında dizi içndeki elemanların toplama işelmini gerçekleştiren fonksiyonumuz
( Our function that performs the collection of elements in the array in the range loop logic )
*/

func SequenceSum(start, end, step int) int {

	total := 0

	for i := start; i <= end; i += step {

		total += i

	}
	return total
}
