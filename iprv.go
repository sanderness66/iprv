// IPRV -- calculate current, power, resistance and/or voltage from any two.
//
// svm 19-MAY-2022 - 13-MAR-2023
//
// I = V/R         R = V/I         V = R*I
// P = I*V         I = P/V         V = P/I
//

package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math"
	"os"
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
			case "i", "I":
				i, _, _ = humanize.ParseSI(xx[1])
			case "p", "P":
				p, _, _ = humanize.ParseSI(xx[1])
			case "v", "V":
				v, _, _ = humanize.ParseSI(xx[1])
			case "r", "R":
				r, _, _ = humanize.ParseSI(xx[1])
			}
		case 2:
			println("bad arg")
			os.Exit(1)
		}
	}

	if p > 0 && v > 0 {
		// println("pv to ir")
		i = p / v
		r = v / i
	} else if v > 0 && r > 0 {
		// println("vr to ip")
		i = v / r
		p = i * v
	} else if i > 0 && v > 0 {
		// println("iv to pr")
		p = i * v
		r = v / i
	} else if p > 0 && i > 0 {
		// println("pi to vr")
		v = p / i
		r = v / i
	} else if r > 0 && i > 0 {
		// println("ri to pv")
		v = r * i
		p = i * v
	} else if r > 0 && p > 0 {
		// println("rp to vi")
		v = math.Sqrt(r * p)
		i = v / r
	} else {
		println("bad arg")
		os.Exit(1)
	}

	prpr("resistance", "R", "Ω", r)
	prpr("voltage", "V", "V", v)
	prpr("current", "I", "A", i)
	prpr("power", "P", "W", p)
}

func prpr(label string, abbrev string, unit string, val float64) {
	vval := humanize.SIWithDigits(val, 2, unit)
	fmt.Printf("%-10s %s = %s\n", label, abbrev, vval)
}
