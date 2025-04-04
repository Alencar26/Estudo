package main

func main() {

	slice := []int{1, 2, 3, 4, 5, 6}
	invertido := inverter(slice)
	invertido2 := inverter2(slice)
	invertido3 := inverter3(slice) //mais performático

	for _, v := range invertido {
		print(v)
	}
	println("")
	for _, v := range invertido2 {
		print(v)
	}
	println("")
	for _, v := range invertido3 {
		print(v)
	}
}

func inverter(arr []int) (arrInvertido []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		arrInvertido = append(arrInvertido, arr[i])
	}
	return
}

func inverter2(arr []int) []int {
	arrAux := make([]int, len(arr))
	i := 0
	for i <= (len(arr) - 1) {
		arrAux[i] = arr[len(arr)-1-i]
		i++
	}
	return arrAux
}

// mantem mesmo array - Sem criar uma lista auxiliar
func inverter3(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
