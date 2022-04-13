package main

import "fmt"

func main() {
	test(1)
	test("23")
	test(map[string]interface{}{"!2": []string{"123", "23423r"}})
}

func test[T any](a T) {
	fmt.Println(a)
}
