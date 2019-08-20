package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNonRepeatingSubStrAd(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("bbbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd(""))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("abcdef"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStrAd("黑化肥挥发发灰会花飞灰化肥会发会发黑会飞花"))

}
