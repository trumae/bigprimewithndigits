package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
)

func main() {
	var n int
	flag.IntVar(&n, "length", 10, "number of digits")

	var nbases int
	flag.IntVar(&nbases, "nbase", 6, "number of pseudorandomly chosen bases")

	flag.Parse()

	banner := ""
	for i := 0; i < n; i++ {
		banner += "9"
	}

	digits := "0123456789"
	for idx := n - 1; idx >= 0; idx-- {
		el := banner[idx]
		if el != '\n' {
			nbanner := banner
			for _, c := range digits {
				runes := []rune(nbanner)
				runes[idx] = c

				bi := new(big.Int)
				bi.SetString(string(runes), 10)
				prime := bi.ProbablyPrime(nbases)
				if prime {
					fmt.Println("* NUMBER *\n")
					fmt.Println(string(runes))
					fmt.Println("\nThe probability of the chosen number being prime is", 1.0-1/math.Pow(4, float64(nbases)))
					return
				}
			}
		}
	}

	fmt.Println("Not found prime to digits:", n)
}
