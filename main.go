package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Books struct {
	Name string `json:"title"`
}

//Prints all the books from the list
func list(Books []Books) {
	for i := 0; i < len(Books); i++ {
		fmt.Printf("Book %d : %s \n", i+1, Books[i].Name)
	}
}

//Checks if the specified book is in the list and prints it
func search(Books []Books, Key string) {
	Lock := true
	for i := 0; i < len(Books); i++ {
		if strings.ToLower(Books[i].Name) == strings.ToLower(Key) {
			fmt.Printf("Book %d : %s \n", i+1, Books[i].Name)
			Lock = false
		} else if i == len(Books)-1 && Lock {
			println("Book is not on the list")
		}
	}

}

func main() {
	args := os.Args[1:]
	var Books []Books
	//Reads JSON file
	file, err := ioutil.ReadFile("books.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Decode JSON file to array
	err = json.Unmarshal(file, &Books)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Check args if args is list and length is 1 it will call list function,
	//otherwise check args and if its search calls search.
	//if none works it, prints 'Please use args'
	if len(args) == 1 && strings.ToLower(args[0]) == "list" {
		list(Books)
	} else if strings.ToLower(args[0]) == "search" {
		search(Books, strings.Join(args[1:], " "))
	} else {
		println("Please use args")
	}
}
