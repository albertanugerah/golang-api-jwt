package main

import "fmt"

func main() {
	palindrome := "abbaa"
	isPalindrome(palindrome)

}



func isPalindrome (a string){

	for i := 0; i<len(a)/2; i++ {
		temp := len(a)-1-i

		if a[i] != a[temp]{
			fmt.Println("Is Not Palindrome")
			return
		}

	}
	fmt.Println("Is Palindrome")

}
