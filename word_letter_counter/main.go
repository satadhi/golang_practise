package main

import (
	"fmt"
	"os"
	"strings"
	"word_letter_counter/counters"
)

func main() {

	file := "input.txt"
	content, error := os.ReadFile(file)
	if error != nil {
		fmt.Println("Error reading file:", error)
		return
	}

	// convert the content to string
	text := string(content)

	// split words and letters from the text
	words := strings.Fields(text)
	letters := strings.ReplaceAll(text, " ", "")


	// Channels for word and letter count results
	wcChan := make(chan map[string]int)
	lcChan := make(chan map[rune]int)

	// GoRoutine to count words
	go func() {
		letterCounts := counters.CountLetters(letters)
		lcChan <- letterCounts
	}()
	
	// GoRoutine to count letters
	go func() {
		wordCounts := counters.CountWords(words)
		wcChan <- wordCounts
	}()


	wordCountResult := <- wcChan
	letterCountResult := <- lcChan
	
	// fmt.Println("Word Count:")
	// for word, count := range wordCountResult {
	// 	fmt.Printf("%s: %d\n", word, count)
		
	// }
	
	// fmt.Println("\nLetter Count: ")
	// for letter, count := range letterCountResult {
	// 	fmt.Printf("%c: %d\n", letter, count)
	// }


	// Open the output file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write word counts to the file
	_, err = outputFile.WriteString("Word Count:\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	for word, count := range wordCountResult {
		_, err := outputFile.WriteString(fmt.Sprintf("%s: %d\n", word, count))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Write letter counts to the file
	_, err = outputFile.WriteString("\nLetter Count:\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	for letter, count := range letterCountResult {
		_, err := outputFile.WriteString(fmt.Sprintf("%c: %d\n", letter, count))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Results saved to output.txt")
}
