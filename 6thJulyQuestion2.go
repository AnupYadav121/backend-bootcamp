package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(id int) {
	fmt.Printf("Rating %d starting\n", id)

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	time.Sleep(time.Duration(rand.Int31n(3000)) * time.Millisecond)
	fmt.Printf("Rating %d done\n", id)
}

func main() {
	rating := 0
	numStudents := 200
	var wg sync.WaitGroup

	for i := 1; i <= numStudents; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			getRating(i)

		}()

		seed := rand.NewSource(time.Now().UnixNano())
		rand := rand.New(seed)
		min := 1
		max := 10
		var ratingRandom int = rand.Intn(((max - min + 1) + min))
		rating += ratingRandom

	}

	wg.Wait()

	fmt.Println(".....")
	fmt.Println("The sum of all ratings of class (out of 10) is :", rating)
	fmt.Println("The average rating of class (out of 10) is :", float64(rating)/float64(numStudents))
}
