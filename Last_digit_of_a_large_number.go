package main

/*
bu fonksiyon iki string sayı verip bunları üssel bir biçimde çarpıp çıkan çarpım sonucunun birler basamağını yazdırmamızı istiyordu
( this function gives us two string numbers and asks us to multiply them exponentially and print the first digit of the product result )
 !! burda dikkat çekmemizi istediği konu sonucun büyük değerler çıkacağı içn Bigint türünde çevrim yapmamızı istemesidir.
*/
import (
	"math/big"
	"strconv"
)

func LastDigit(n1, n2 string) int {

	if n2 == "0" { // Üs 0 ise, sonuç 1'dir.
		return 1
	}

	bigN1 := new(big.Int)                // tanımlama yapıyoruz
	bigN1, ok := bigN1.SetString(n1, 10) // burada da string türündeki bir veriyi BigInt türüne çeviriyoruz
	if !ok {
		return -1
	}

	bigN2 := new(big.Int)
	bigN2, ok = bigN2.SetString(n2, 10)
	if !ok {
		return -1
	}

	result := new(big.Int)
	result.Exp(bigN1, bigN2, big.NewInt(10)) // EXP fonksiyonu mod alma işlemi yapar ve ardından 3.paremetreye göre mod işlemini uygular
	//result.Exp(bigN1, bigN2, nil) bu şekilde 3.paremtre nil olursa mod işlemi uygulanmaz sadece üs alma işelmi uygulanır

	lastDigit, _ := strconv.Atoi(result.String())

	return lastDigit
}
