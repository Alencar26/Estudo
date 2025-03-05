package main

import "fmt"

func main() {

	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(ss)       //all
	fmt.Println(ss[0])    //row
	fmt.Println(ss[1])    //row
	fmt.Println(ss[2])    //row
	fmt.Println(ss[0][2]) //value
	fmt.Println(ss[1][0]) //value
	fmt.Println(ss[2][1]) //value

	//multiniveis
	ssss := [][][][]int{
		{
			{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 0},
			},
			{
				{3, 4, 5, 6, 7},
			},
		},
		{
			{
				{4, 4, 4, 4, 4, 4},
				{4, 4, 4, 4, 4, 4},
			},
		},
		{
			{
				{3},
			},
			{
				{2, 2, 2, 2, 2, 2},
				{5, 5, 5, 5, 5},
			},
			{
				{4, 8},
			},
		},
	}

	fmt.Println(ssss[:])          //todo o slice
	fmt.Println(ssss[0][1][0][3]) //obtendo valor específico
}
