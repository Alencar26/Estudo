package main

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	result := removeParesDoSlice(arr)
	for _, v := range result {
		print(v)
	}
}

func removeParesDoSlice(arr []int) (oddNubers []int) {
	for _, v := range arr {
		if v%2 != 0 {
			oddNubers = append(oddNubers, v)
		}
	}
	return
}
