package main

import "fmt"

func main() {
	var longitud int
	var suma int
	fmt.Scan(&longitud)
	s := make([]int, longitud)
	for i := 0; i < len(s); i++ {
		fmt.Scan(&s[i])
	}
	for _, v := range s {
		suma += v
	}
	fmt.Println(suma)
}
