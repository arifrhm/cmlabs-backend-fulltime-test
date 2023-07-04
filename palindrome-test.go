package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	// Menghapus spasi dan mengubah huruf menjadi lowercase
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	// Membandingkan string dengan versi terbaliknya
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("SAIPPUAKIVIKAUPPIAS"))  // true
	fmt.Println(isPalindrome("Aibohphobia"))          // true
	fmt.Println(isPalindrome("Anna"))                 // true
	fmt.Println(isPalindrome("Civic"))                // true
	fmt.Println(isPalindrome("My gym"))               // true
	fmt.Println(isPalindrome("No lemon, no melon"))   // true

	fmt.Println(isPalindrome("Hello"))                // false
	fmt.Println(isPalindrome("OpenAI"))               // false
}
