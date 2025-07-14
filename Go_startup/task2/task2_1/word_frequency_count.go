package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)


func HelperFunction(input string) map[string]int {
    
    frequency := make(map[string]int)
    var str_builder strings.Builder

    for _,word := range input{
        if unicode.IsLetter(word) || unicode.IsSpace(word){
            str_builder.WriteRune(unicode.ToLower(word))
        }
    }
    words := strings.Fields(str_builder.String())
    for _, word := range words {
        frequency[word]++
    }
    
    return frequency
}
func main(){

    fmt.Println("Enter your sentence?")
    
    reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
    count := HelperFunction(word)

    for wrd , cnt := range count{
        fmt.Println(wrd , cnt)
    }
}