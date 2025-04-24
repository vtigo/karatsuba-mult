package main

import (
	"fmt"
	"math/big"
)

func main() {
	num1 := new(big.Int)
	num1.SetString("3141592653589793238462643383279502884197169399375105820974944592", 10)
	num2 := new(big.Int)
	num2.SetString("2718281828459045235360287471352662497757247093699959574966967627", 10)
	
	result := multiply(num1, num2)
	fmt.Println(result.String())
}

func multiply(x *big.Int, y *big.Int) *big.Int {
	// Base case for recursion
	ten := big.NewInt(10)
	if x.Cmp(big.NewInt(10)) < 0 || y.Cmp(big.NewInt(10)) < 0 {
		return new(big.Int).Mul(x, y)
	}
	
	// Calculate the size of the numbers
	xLen := intLen(x)
	yLen := intLen(y)
	n := max(yLen, xLen)
	
	m := (n + 1) / 2 // Ceiling of n/2
	
	// Calculate 10^m
	pow10m := new(big.Int).Exp(ten, big.NewInt(int64(m)), nil)
	
	// Split the numbers
	a, b := new(big.Int), new(big.Int)
	a.DivMod(x, pow10m, b) // a = x / 10^m, b = x % 10^m
	
	c, d := new(big.Int), new(big.Int)
	c.DivMod(y, pow10m, d) // c = y / 10^m, d = y % 10^m
	
	// Recursive steps
	ac := multiply(a, c)
	bd := multiply(b, d)
	
	// Calculate (a+b)*(c+d)
	aPlusB := new(big.Int).Add(a, b)
	cPlusD := new(big.Int).Add(c, d)
	abcd := multiply(aPlusB, cPlusD)
	
	// abcd = (a+b)*(c+d) - ac - bd
	abcd.Sub(abcd, ac)
	abcd.Sub(abcd, bd)
	
	// Calculate 10^(2*m)
	pow10_2m := new(big.Int).Exp(ten, big.NewInt(int64(2*m)), nil)
	
	// Calculate ac * 10^(2*m)
	acTerm := new(big.Int).Mul(ac, pow10_2m)
	
	// Calculate abcd * 10^m
	abcdTerm := new(big.Int).Mul(abcd, pow10m)
	
	// Calculate the final result: ac*10^(2*m) + abcd*10^m + bd
	result := new(big.Int).Add(acTerm, abcdTerm)
	result.Add(result, bd)
	
	return result
}

func intLen(n *big.Int) int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return 1
	}
	
	temp := new(big.Int).Set(n)
	
	count := 0
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	
	for temp.Cmp(zero) > 0 {
		temp.Div(temp, ten)
		count++
	}
	
	return count
}
