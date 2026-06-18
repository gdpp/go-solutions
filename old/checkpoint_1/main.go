package main

import "fmt"

func main() {

	summonerName := "Nerion"
	pi := 3.1416
	lives := 3

	fmt.Println(summonerName)
	fmt.Println(pi)
	fmt.Println(lives)

	hits := []int{2, 3, 6, 1, 8}

	suma := 0

	for x := 0; x < len(hits); x++ {
		suma = suma + hits[x]
	}

	fmt.Println(suma)

	sayHello("Gustavo")

	sum(10, 15)

	nums := []int{10, 15, 34, 12, 8, 76}
	avg := average(nums)

	fmt.Println(avg)

	countDown(100)

	showItems([]string{"Hola", "Mundo", "3IAtlas", "Aliens", "Apocalypse", "Escape", "Fate"})

	evenOrOdd([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

}
