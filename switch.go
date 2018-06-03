package main

import (
	"time"
	"fmt"
)

func main() {

	weekday := time.Now().Weekday()

	switch weekday {
	case time.Saturday:
		fmt.Println("today is saturday")

	case time.Sunday:
		fmt.Println("today is sunday")

	default:
		fmt.Println("today is busy day")
	}

	second := time.Now().Second()
	switch {

	case second >= 0 && second < 30:
		fmt.Println("上半分钟")

	default:
		fmt.Println("下班分钟")
	}

	whatType := func(value interface{}) {
		switch valueType := value.(type) {
		case bool:
			fmt.Println("this is boolean")

		case string:
			fmt.Println("this is string")

		case int:
			fmt.Println("this is int")

		default:
			fmt.Printf("Wonna know the the fuck type: %T\n", valueType)
		}
	}

	whatType(true)
	whatType(12)
	whatType(1.2)
	whatType("string")
}
