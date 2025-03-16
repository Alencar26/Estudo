package main

import "fmt"

func main()  {
  //CALLBACK é passar um função coomo parâmetro para outra função.
  t := somentePares(soma, []int{2,3,4,1,4,22,5,6,88,7,77,97,91,3,25,64}...)
  fmt.Println(t)
}


func soma(x ...int) int  {
  n := 0
  for _, v := range x {
    n += v
  }
  return n
}

func somentePares(f func(x ...int) int, y ...int) int  {
  var slice []int
  for _,v := range y {
    if v % 2 == 0 {
      slice = append(slice, v)
    }
  }
  total := f(slice...)
  return total
}
