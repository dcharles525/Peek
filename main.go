package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

// Define the character set
const charSet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678
	9 ,.!?`
const seed = 1
const bookLength = 32000

func generateRandomString(wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	b := make([]byte, bookLength)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	results <- string(b)
}

func seekBook(bookNumber int) string {
	var wg sync.WaitGroup
	results := make(chan string, 1)

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < bookNumber; i++ {
		wg.Add(1)
		go generateRandomString(&wg, results)
	}
	return <-results
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <book>")
		return
	}

	numBooks, numBooksError := strconv.Atoi(os.Args[1])
	rand.Seed(seed)

	if numBooksError != nil {
		fmt.Println("Usage: Enter a valid integer value for the book number")
	}

	fmt.Println(seekBook(numBooks))
}
