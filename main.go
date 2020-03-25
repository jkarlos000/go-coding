package main

import "fmt"

func isBalanced(p string) bool {
	pairs := map[rune]rune{'(':')', '{': '}', '[': ']'} //Rune - alias Int32
	stack := []rune{}

	for _, c := range p {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) < 1 {
				return false
			}
			if c == pairs[stack[len(stack) - 1]] {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	// Diccionarios (asi se conoce en Python) de Go lang (MAP)
	parentheses := []string{")(", "[(])", "(()){}[]", "())[]{}", "()[]{}(([])){[()][]}"}
	for _, p := range parentheses{
		fmt.Printf("%s: %t\n", p, isBalanced(p))
	}
}
