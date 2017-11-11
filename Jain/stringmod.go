package main
// stringmod is a way to change a string to make another striung with the changes.

//A string is a sequence of Unicode character
//String is an immutable type variable
//to change a string once its created, 
//you need to first convert a slice into a rune
//then do the changes and in the end convert it back to string

import "fmt" 

func main() {
	s := "hello, World!"
	r := []rune(s)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
}
