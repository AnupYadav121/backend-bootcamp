package main

import "fmt"

func getMap(channel chan map[rune]int, val string) {
	newMap := make(map[rune]int)

	for _, chrc := range val {
		_, booll := newMap[chrc]
		if booll == false {
			newMap[chrc] = 1
		} else {
			newMap[chrc] += 1
		}
	}
	channel <- newMap

}

func main() {
	global := make(map[string]int)
	strArr := [5]string{"quick", "brown", "fox", "lazy", "dog"}

	channel := make(chan map[rune]int, 5)

	for _, elem := range strArr {

		go func(elem string) {
			getMap(channel, elem)
		}(elem)

		getMap := <-channel

		for key, val := range getMap {
			_, booll := global[string(key)]
			if booll == false {
				global[string(key)] = val
			} else {
				global[string(key)] += val
			}
		}
	}

	fmt.Println(global)

}
