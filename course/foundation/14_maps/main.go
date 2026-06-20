package main

import "fmt"

func main() {

	// map[keyType]valueType
	ages := map[string]int{
		"alice":   30,
		"bob":     25,
		"charlie": 35,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println(ages["alice"])

	var scores map[string]int

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)

	scores["math"] = 90
	scores["english"] = 85
	scores["science"] = 95

	fmt.Println(scores)

	delete(scores, "math")

	fmt.Println(scores)

	points := map[string]int{
		"a": 10,
		"b": 0, // valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	fmt.Println("b", valB, okB)

	valC, okC := points["c"]
	fmt.Println("c", valC, okC)

	if val, ok := points["c"]; ok {
		fmt.Println("c", val)
	} else {
		fmt.Println("c does not exist")
	}

	total := 0

	for item, score := range scores {
		fmt.Println(item, score)
		total = total + score
	}

	fmt.Println("Total:", total)

}
