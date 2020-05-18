package main

import (
	"fmt"
	"os"

	"github.com/blck-snwmn/booksearch"
	"golang.org/x/xerrors"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Println("error!")
		return
	}
	sub := os.Args[1]
	target := os.Args[2]
	var err error
	switch sub {
	case "search":
		err = booksearch.Search(target)
	case "register":
		err = booksearch.Register(target)
	default:
		err = xerrors.New("no command")
	}
	if err != nil {
		fmt.Println(err)
	}
}
