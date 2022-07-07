package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(id int, chnl chan int, wg *sync.WaitGroup) {
	fmt.Printf("Rating %d starting\n", id)
	seed := rand.NewSource(time.Now().UnixNano())
	randy := rand.New(seed)
	time.Sleep(time.Duration(randy.Int31n(3000)) * time.Millisecond)
	fmt.Printf("Rating %d done\n", id)

	min := 1
	max := 10
	var ratingRandom int = randy.Intn(((max - min + 1) + min))
	chnl <- ratingRandom
	wg.Done()
}

func getUpdate(val *int, mych chan int, n int) {
	for i := 0; i < n; i++ {
		ratingGiven := <-mych
		*val += ratingGiven
	}
}

func main() {
	rating := 0
	numStudents := 50
	var wg sync.WaitGroup
	mych := make(chan int, 5)
	go getUpdate(&rating, mych, numStudents)

	for i := 1; i <= numStudents; i++ {
		wg.Add(1)
		j := i

		go getRating(j, mych, &wg)

	}

	wg.Wait()

	fmt.Println(".....")
	fmt.Println("The sum of all ratings of class (out of 10) is :", rating)
	fmt.Println("The average rating of class (out of 10) is :", float64(rating)/float64(numStudents))
}
