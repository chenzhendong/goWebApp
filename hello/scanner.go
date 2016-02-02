package main
import (
	"fmt"
	"os"
	"bufio"
	"text/scanner"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("/Users/a4d98zz/src/golang/src/goWebApp/hello/input.txt")
	check(err)
	reader := bufio.NewReader(file)
	var s scanner.Scanner
	s.Init(reader)
	s.TokenText()

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}

}