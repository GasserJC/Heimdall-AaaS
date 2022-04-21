package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func reportPerformance() {
	//reportPerformanceAtSize(100)
	//reportPerformanceAtSize(10000)
	reportPerformanceAtSize(1000000)
	//reportPerformanceAtSize(100000000)
}

func reportPerformanceAtSize(size int) {
	os.RemoveAll("./data/performanceTest")
	usernames := addRandomValuesToSlice(size, 8)
	passwords := addRandomValuesToSlice(size, 16)

	now := time.Now()
	for i := 0; i <= size; i++ {
		AddUser("performanceTest", usernames[i], passwords[i])
	}
	then := time.Now()
	fmt.Print("AddUser at ")
	fmt.Print(size)
	fmt.Print(" took ")
	fmt.Println(then.Sub(now) / time.Duration(size))

	now = time.Now()
	for i := 0; i <= size; i++ {
		AuthUser("performanceTest", usernames[i], passwords[i])
	}
	then = time.Now()
	fmt.Print("AuthUser at ")
	fmt.Print(size)
	fmt.Print(" took ")
	fmt.Println(then.Sub(now) / time.Duration(size))
}

func addRandomValuesToSlice(amountOfValues int, lengthOfValues int) []string {
	var randomsSlice []string
	for i := 0; i <= amountOfValues; i++ {
		randomsSlice = append(randomsSlice, randomString(lengthOfValues, i))
	}
	return randomsSlice
}

func randomString(n int, modifier int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
