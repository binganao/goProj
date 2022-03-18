package main

import (
	"fmt"
	"strconv"
	"strings"
)

func prepend(slice []string, ele string) []string {
	return append([]string{ele}, slice...)
}

func main() {
	var n string
	var sum int
	fmt.Println(strings.Join([]string{"1", "2"}, "--"))
	fmt.Scanf("%s", &n)
	for i := 0; i < len(n); i++ {
		place, _ := strconv.Atoi(string(n[i]))
		sum += place
	}
	res := []string{""}
	py := []string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}
	for ; sum > 0; sum /= 10 {
		res = prepend(res, py[sum%10])
	}
	for i := 0; i < len(res); i++ {
		if i != 0 && i != len(res)-1 {
			fmt.Print(" ")
		}
		fmt.Print(res[i])
	}
}
