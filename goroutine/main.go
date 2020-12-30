package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"jinseo", "kay"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting for messages")

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for: ", i)
		resultChan := <-c
		fmt.Println("Received this message: ", resultChan)
	}

	// result := <-c
	// fmt.Println(result)
}

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
