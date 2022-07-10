package main

import "fmt"

func getMap(channel chan map[rune]int, val string) {
	newMap := make(map[rune]int)
	for _, chrc := range val {
		newMap[chrc]++
	}
	channel <- newMap

}

func main() {
	global := make(map[string]int)
	strArr := [5]string{"quick", "brown", "fox", "lazy", "dog"}
	channel := make(chan map[rune]int, 5)

	for _, elem := range strArr {
		go getMap(channel, elem)
		getMap := <-channel

		for key, val := range getMap {
			global[string(key)] += val
		}
	}

	fmt.Println(global)

}
