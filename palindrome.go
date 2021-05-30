package main

import "fmt"

func main() {
	palindrome := "abbaa"
	isPalindrome(palindrome)

	if isPalindromeRecursive(palindrome) == true{
		fmt.Println("Is Palindrome")
	}else{
		fmt.Println("Is Not Palindrome")
	}


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

func isPalindromeRecursive(text string) bool {

	if len(text) <= 1{
		return true
	}else if len(text) == 2 && text[0] == text[1] {
		return  true
	}else if text[0] != text[len(text)-1] {
		return false
	}

	return isPalindromeRecursive(text[1 : len(text)-1])
}
