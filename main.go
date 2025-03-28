package main

import (
	"fmt"
  "sort"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

const charSet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ,.!?`
const seed = 1
const bookLength = 32000

func generateRandomString(wg *sync.WaitGroup, f int, books chan<- map[int]string) {
	defer wg.Done()
  randSeeded := rand.New(rand.NewSource(int64(f)))
	b := make([]byte, bookLength + 100)
  for i := 0; i < bookLength; i++ {
		b[i] = charSet[randSeeded.Intn(len(charSet))]
	}

	book := make(map[int]string)
	book[f] = string(b)
	books <- book
}

func seekBook(bookNumber int) string {
	var wg sync.WaitGroup
	books := make(chan map[int]string, bookNumber)

	go func() {
		wg.Wait()
		close(books)
	}()

	for i := 0; i < bookNumber; i++ {
		wg.Add(1)
		go generateRandomString(&wg, i, books)
	}

	finalResults := make(map[int]string)
	for book := range books {
		for k, v := range book {
			finalResults[k] = v
		}
	}

  keys := make([]int, 0, len(finalResults))
	for k := range finalResults {
		keys = append(keys, k)
	}
  sort.Ints(keys)

	return finalResults[keys[len(keys)-1]]
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <book>")
		return
	}

	numBooks, numBooksError := strconv.Atoi(os.Args[1])

	if numBooksError != nil {
		fmt.Println("Usage: Enter a valid integer value for the book number")
	}

	fmt.Println(seekBook(numBooks))
}
