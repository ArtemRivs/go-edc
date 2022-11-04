package main

import "fmt"

func main() {

	words := []string{"hello", "world", "hello", "world", "world"}
	wm := make(map[string]int, 0)
	for _, word := range words {
		if val, ok := wm[word]; ok {
			wm[word] = val + 1
			continue
		}
		wm[word] = 1
	}
	fmt.Println("words: ", words)
	fmt.Println("wm: ", wm)

	var maxWord string
	var maxValue int
	for key, val := range wm {
		if val > maxValue {
			maxValue = val
			maxWord = key
		}
	}

	fmt.Println("maxWord: ", maxWord)

}
