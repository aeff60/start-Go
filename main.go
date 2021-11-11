package main

import(
	"os"
	"strconv"
)

func sum(lhs string, rhs string) int{
	int1, _ := strconv.Atoi(lhs)
	int2, _ := strconv.Atoi(rhs)
	return int1 + int2
}

func main(){
	no1, _ := strconv.Atoi(os.Args[1])
	no2, _ := strconv.Atoi(os.Args[2])

	println("Sum", no1+no2)
	println("Summa", sum(os.Args[1], os.Args[2]))
}
