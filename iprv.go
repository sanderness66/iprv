// IPRV -- calculate current, power, resistance and/or voltage from any two.
//
// svm 19-MAY-2022 - 22-MAY-2022
//
// I = V/R         R = V/I         V = R*I
// P = I*V         I = P/V         V = P/I
//

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var i, p, v, r float64

	args := os.Args[1:]
	for ac, av := range args {
		switch ac {
		case 0, 1:
			// we might want to implement some sort of
			// sanity checking here some day.
			xx := strings.Split(av, "=")
			switch xx[0] {
			case "i", "I", "a", "A":
				i, _ = strconv.ParseFloat(xx[1], 32)
			case "p", "P":
				p, _ = strconv.ParseFloat(xx[1], 32)
			case "v", "V":
				v, _ = strconv.ParseFloat(xx[1], 32)
			case "r", "R":
				r, _ = strconv.ParseFloat(xx[1], 32)
			}
		case 2:
			println("bad arg")
			os.Exit(1)
		}
	}

	if p > 0 && v > 0 {
		println("pv to ir")
		i = p / v
		r = v / i
	} else if v > 0 && r > 0 {
		println("vr to ip")
		i = v / r
		p = i * v
	} else if i > 0 && v > 0 {
		println("iv to pr")
		p = i * v
		r = v / i
	} else if p > 0 && i > 0 {
		println("pi to vr")
		v = p / i
		r = v / i
	} else if r > 0 && i > 0 {
		println("ri to pv")
		v = r * i
		p = i * v
	} else if r > 0 && p > 0 {
		println("rp to vi")
		v = math.Sqrt(r * p)
		i = v / r
	} else {
		println("bad arg")
		os.Exit(1)
	}

	fmt.Printf("resistance: R = %1.3g\n", r)
	fmt.Printf("voltage:    V = %1.3g\n", v)
	fmt.Printf("current:    I = %1.3g\n", i)
	fmt.Printf("power:      P = %1.3g\n", p)
}
