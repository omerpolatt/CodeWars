package main

/*
Verilen stringi ikili ikili ayıran fonksiyonumuz
( Our function to split the given string binary by binary )
*/

func Solution(str string) []string {
	result := []string{}

	word := ""

	for _, i := range str {

		word += string(i)

		if len(word) == 2 {
			result = append(result, word)
			word = ""
		}

	}

	if len(word) == 1 {
		word += "_"
		result = append(result, word)
		word = ""

	}

	return result

}
