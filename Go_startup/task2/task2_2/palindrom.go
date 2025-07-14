package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindom(words string) bool {
	runes := []rune(words)
	right := len(runes) - 1
	left := 0
	

	for left <= right{
		if runes[left] != runes[right]{
			return false
			
		}
		right --
		left ++

	}

	return true
}
func main(){

	fmt.Println("Enter Your Sentence?")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')

	var str_builder strings.Builder

	for _ , words := range word{

		if unicode.IsLetter(words) {
			str_builder.WriteRune(unicode.ToLower(words))
		}
	}

	normalize := str_builder.String()
	fmt.Println("IS palindrome:")
	fmt.Println(palindom(normalize))


}