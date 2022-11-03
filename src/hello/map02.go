package main


func UniqueUserIDs(userIDs []int64) []int64{

	unic := map[int64]bool{}
	newSlice :=[]int64{}
	for i := 0; i < len(userIDs); i++ {
		if unic[userIDs[i]] == true {
			continue
		}
		unic[userIDs[i]] = true
		newSlice = append(newSlice, userIDs[i])
	}

	return newSlice

}
