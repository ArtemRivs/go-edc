package main


func UniqueUserIDs(userIDs []int64) []int64{

	unic := map[int64]bool{}
	newSlice :=[]int64{}
	for c, id := range userIDs {
		if map[id] !=false {
			continue
		}
		unic[id] = true
		newSlice = append(newSlice, map[id])
	}

	return newSlice

}


