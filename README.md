## About

This is a basic CLI application to check the books list.

## Commands & How to use
First, install go.
Download project as zip from Github.
Extract the project to your computer.
Optionally you can change books.json.
Open your terminal in the project folder.
### list command
```
go run main.go list
```
This command will return all books from the booklist.

### search command 
```
go run main.go search <bookName>
```
This command will return the specified book from the book list if it's available.