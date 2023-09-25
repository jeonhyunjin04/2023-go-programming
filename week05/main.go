package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	repeaceWords := strings.NewReplacer("?", "o")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)

}
