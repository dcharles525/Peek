package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
)

const charSet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678
	9 ,.!?`
const seed = 1
const bookLength = 32 //32000

func generateRandomString(wg *sync.WaitGroup, i int, books chan<- map[int]string, r *rand.Rand) {
	defer wg.Done()
	b := make([]byte, bookLength)
	for i := range b {
		b[i] = charSet[r.Intn(len(charSet))]
	}

	book := make(map[int]string)
	book[i] = string(b)
	books <- book
}

func seekBook(bookNumber int) string {
	var wg sync.WaitGroup
	books := make(chan map[int]string, bookNumber)
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < bookNumber; i++ {
		wg.Add(1)
		go generateRandomString(&wg, i, books, r)
	}

	go func() {
		wg.Wait()
		close(books)
	}()

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
