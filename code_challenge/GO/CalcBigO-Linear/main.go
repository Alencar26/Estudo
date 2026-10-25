package main

func main() {
	n := 10
	result := calculate(n)
	println("The sum of the first", n, "numbers is:", result)
}

func calculate(n int) int {
	// O(n) complexidade de tempo
	sum := 0                  // 1
	for i := 0; i <= n; i++ { //1 + 3 * (n + 1) + 3n = (6n + 2)
		sum += i // 1 + 1 + 1 + 1 = 4*n = 4n
	}
	return sum // 1 + 1 = 2
}

//Total de 10n + 5
//Complexidade de tempo: O(n)
