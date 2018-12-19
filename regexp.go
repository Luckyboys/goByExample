package main

import (
	"fmt"
	"regexp"
)

const RegularExpression = "^[a-zA-Z\\p{Han}]+"
const RegularExpression2 = "^https?://"

func main() {

	str1 := "hello中国hello世界和平hi好"
	str2 := "中国hello世界和平hi好"
	str3 := "+中国hello世界和平hi好"
	str4 := "_dd"
	str5 := "http://www"
	str6 := "https://www"
	str7 := " https://www"

	fmt.Printf("RegularExpression: %s\n", RegularExpression)
	fmt.Printf("%s: %v\n", str1, match(RegularExpression, str1))
	fmt.Printf("%s: %v\n", str2, match(RegularExpression, str2))
	fmt.Printf("%s: %v\n", str3, match(RegularExpression, str3))
	fmt.Printf("%s: %v\n", str4, match(RegularExpression, str4))
	fmt.Printf("%s: %v\n", str5, match(RegularExpression, str5))
	fmt.Printf("%s: %v\n", str6, match(RegularExpression, str6))
	fmt.Printf("%s: %v\n", str7, match(RegularExpression, str7))

	fmt.Println()

	fmt.Printf("RegularExpression: %s\n", RegularExpression2)
	fmt.Printf("%s: %v\n", str1, match(RegularExpression2, str1))
	fmt.Printf("%s: %v\n", str2, match(RegularExpression2, str2))
	fmt.Printf("%s: %v\n", str3, match(RegularExpression2, str3))
	fmt.Printf("%s: %v\n", str4, match(RegularExpression2, str4))
	fmt.Printf("%s: %v\n", str5, match(RegularExpression2, str5))
	fmt.Printf("%s: %v\n", str6, match(RegularExpression2, str6))
	fmt.Printf("%s: %v\n", str7, match(RegularExpression2, str7))
}

func match(regularExpression, str string) bool {

	re, err := regexp.Compile(regularExpression)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return re.MatchString(string(str))
}
