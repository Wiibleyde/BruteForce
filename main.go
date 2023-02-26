package main

import (
	"fmt"
	"time"
	"unicode"
)

func upperCase(size int, prefix string, useSymbol bool, channel chan<- string) {
	letters := "abcdefghijklmnopqrstuvwxyz"
	generateCombinations(prefix, letters, size, useSymbol, channel)
}

func lowerCase(size int, prefix string, useSymbol bool, channel chan<- string) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	generateCombinations(prefix, letters, size, useSymbol, channel)
}

func generateCombinations(prefix string, letters string, size int, useSymbol bool, channel chan<- string) {
	if size == 0 {
		channel <- prefix
		return
	}

	for i := 0; i < len(letters); i++ {
		newPrefix := prefix + string(letters[i])
		generateCombinations(newPrefix, letters, size-1, useSymbol, channel)
	}

	if useSymbol {
		symbols := "!@#$%^&*()_+"
		for i := 0; i < len(symbols); i++ {
			newPrefix := prefix + string(symbols[i])
			generateCombinations(newPrefix, letters, size-1, false, channel)
		}
	}

	// Add the same letter in uppercase
	for i := 0; i < len(letters); i++ {
		if unicode.IsLower(rune(letters[i])) {
			newPrefix := prefix + string(unicode.ToUpper(rune(letters[i])))
			generateCombinations(newPrefix, letters, size-1, useSymbol, channel)
		}
	}
}

func testMdp(password string, tentative string) bool {
	if password == tentative {
		return true
	}
	return false
}

func main() {
	password := "nath"
	passwordSize := len(password)
	startTime := time.Now()
	channel := make(chan string)
	go upperCase(passwordSize, "", false, channel)
	go lowerCase(passwordSize, "", false, channel)
	for {
		tentative := <-channel
		fmt.Println("Tentative: ", tentative)
		if testMdp(password, tentative) {
			fmt.Println("Password found: ", tentative)
			break
		}
	}
	endTime := time.Now()
	fmt.Println("Time: ", endTime.Sub(startTime))
}
