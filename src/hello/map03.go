package solution

func MostPopularWord(words []string) string {

	// wm := map[string]int{}
	wm := make(map[string]int,0)
	for _, word := range words {
		if val, ok := wm[word]; ok {
			wm[word] = val + 1
		}
	}

	var maxWord string
	var maxValue int 
	for key, val := range wm {
		if val > maxValue {
			maxValue = val
			maxWord = key
		}
	}
	
	return maxWord
}
