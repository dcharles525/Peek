package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const charSet = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ,.!?`
const seed = 1
const bookLength = 32000

func generateRandomString(bookPage int) []byte {
	randSeeded := rand.New(rand.NewSource(int64(bookPage)))
	book := make([]byte, bookLength)

	for i := 0; i < bookLength; i++ {
		book[i] = charSet[randSeeded.Intn(len(charSet))]
	}

	return book
}

func seekBook(args []string) []byte {
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <book>")
		return []byte{}
	}

	numBooks, numBooksError := strconv.Atoi(args[1])

	if numBooksError != nil {
		fmt.Println("Usage: Enter a valid integer value for the book number")
		return []byte{}
	}

	return generateRandomString(numBooks)
}

func main() {
	fmt.Println(string(seekBook(os.Args)))
}
