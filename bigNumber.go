package main


import (
		"fmt"
		"math/big"
		)


func main(){

	x, y := big.NewFloat(10), big.NewFloat(2)
	z := big.NewFloat(20).Quo(x, y)
	fmt.Println(x,y,z);
}