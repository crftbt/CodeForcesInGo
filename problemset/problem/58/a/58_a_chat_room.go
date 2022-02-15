/* https://codeforces.com/problemset/problem/58/A
A. Chat room
time limit per test1 second
memory limit per test256 megabytes
inputstandard input
outputstandard output
Vasya has recently learned to type and log on to the Internet. He immediately entered a chat room and decided to say hello to everybody. Vasya typed the word s. It is considered that Vasya managed to say hello if several letters can be deleted from the typed word so that it resulted in the word "hello". For example, if Vasya types the word "ahhellllloou", it will be considered that he said hello, and if he types "hlelo", it will be considered that Vasya got misunderstood and he didn't manage to say hello. Determine whether Vasya managed to say hello by the given word s.

Input
The first and only line contains the word s, which Vasya typed. This word consisits of small Latin letters, its length is no less that 1 and no more than 100 letters.

Output
If Vasya managed to say hello, print "YES", otherwise print "NO".

Examples
inputCopy
ahhellllloou
outputCopy
YES
inputCopy
hlelo
outputCopy
NO
*/

package main

import "fmt"

func input_calc(input_string string) string {
	var hello_index int
	var yn string
	hello := []rune("hello")
	fmt.Scanln(&input_string)
	for _, letter := range input_string {
		if letter == hello[hello_index] {
			hello_index++
			if hello_index == 5 {
				yn = "YES"
				return yn
			}
		}
	}
	yn = "NO"
	return yn
}
func main() {
	fmt.Println(input_calc(""))
}
