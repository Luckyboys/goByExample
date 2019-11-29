package main

import "fmt"

func main() {
	str := make([]string, 3)
	fmt.Println(str)

	str[0] = "e"
	str[1] = "f"
	str[2] = "g"

	fmt.Println(str)
	fmt.Println(str[2])
	fmt.Println(len(str))

	str = append(str, "a", "b")
	str = append(str, "c")
	fmt.Println(str)

	str2 := make([]string, len(str))
	copy(str2, str)

	fmt.Println("str2 = ", str2)
	sli := str[2:5]
	fmt.Println(sli)

	fmt.Println(str[3:])
	fmt.Println(str[:3])

	str3 := []string{"ww", "dd", "ccc"}
	fmt.Println(str3)

	fmt.Println("9x9 table")

	table := make([][]int, 9)
	for x := 1; x <= 9; x++ {

		table[x-1] = make([]int, x)
		for y := 1; y <= x; y++ {
			table[x-1][y-1] = x * y
		}
	}

	printTable(table)

	var keys []int
	fmt.Println(keys)
	fmt.Println("cap: ", cap(keys))
	keys = append(keys, 1)
	keys = append(keys, 2)
	keys = append(keys, 3)
	fmt.Println(keys)
}

/**
 * 打印表格内容
 * @param [][]int table 一个整形的表格，用来展示用
 */
func printTable(table [][]int) {

	fmt.Println("[")
	for x := range table {

		fmt.Printf("\t[")
		notFirstElement := false
		for y := range table[x] {

			if notFirstElement {
				fmt.Printf("\t")
			}

			fmt.Printf("%d", table[x][y])

			notFirstElement = true
		}

		fmt.Printf("]\n")
	}
	fmt.Println("]")
}
