package main

func main() {
	printValues(3)
}

func printValues(n int) {
	for i := 1; i <= n; i++ { // 2 + 6n
		for j := 1; j <= n; j++ { // 2n + (6n)²
			println(i, j) // 3n²
		}
		println("Fim") //n
	}
	println("Fora do loop") //1
}

// O(n²) = 3 + 9n + 9n²
