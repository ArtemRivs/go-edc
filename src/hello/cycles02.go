package solution

func Map(strs []string, mapFunc func(s string) string) []string {

	slice := make([]string, len(strs)) 
	for i, str := range strs {
    	slice[i] = mapFunc(str)
	}
	return slice

}
