package main

import (
	"aoc21/aocutil"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var w float64
var x float64
var y float64
var z float64

func getdest(arg string) *float64 {
	switch arg {
	case "w":
		return &w
	case "x":
		return &x
	case "y":
		return &y
	case "z":
		return &z
	}

	return nil
}

func getmatharg(arg string) float64 {
	parm := getdest(arg)
	var result float64

	if parm == nil {
		a, _ := strconv.Atoi(arg)
		result = float64(a)
	} else {
		result = *parm
	}

	return result
}

func inp(arg string, inqueue *[]int64) {
	var dest *float64

	dest = getdest(arg)
	*dest = float64((*inqueue)[0])
	*inqueue = (*inqueue)[1:]
}

func add(arg1, arg2 string) {
	dest := getdest(arg1)

	arg := getmatharg(arg2)

	*dest = *dest + arg
}

func mul(arg1, arg2 string) {
	dest := getdest(arg1)
	arg := getmatharg(arg2)

	*dest = *dest * arg
}

func div(arg1, arg2 string) {
	dest := getdest(arg1)
	arg := getmatharg(arg2)

	if arg < 0 {
		log.Fatal("Div: Divide by zero!")
	}
	*dest = math.Trunc(*dest / arg)
}

func mod(arg1, arg2 string) {
	dest := getdest(arg1)
	arg := getmatharg(arg2)

	if *dest < 0 || arg <= 0 {
		log.Fatal("Mod: divide by zero!")
	}

	*dest = *dest - math.Trunc(*dest/arg)
}

func eql(arg1, arg2 string) {
	dest := getdest(arg1)
	arg := getmatharg(arg2)

	if *dest == arg {
		*dest = 1
	} else {
		*dest = 0
	}
}

func main() {

	input := "input24.txt"

	lines := aocutil.LoadStringArray(input)

	var inqueue []int64

	// Now start working my way down from the highest number to find the largest
	// one.

	var n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, n14 int64

	for n14 = 9; n14 > 0; n14-- {
		fmt.Printf("n14=%d\n", 9)
		for n13 = 9; n13 > 0; n13-- {
			for n12 = 9; n12 > 0; n12-- {
				for n11 = 9; n11 > 0; n11-- {
					for n10 = 9; n10 > 0; n10-- {
						for n9 = 9; n9 > 0; n9-- {
							for n8 = 9; n8 > 0; n8-- {
								for n7 = 9; n7 > 0; n7-- {
									for n6 = 9; n6 > 0; n6-- {
										for n5 = 9; n5 > 0; n5-- {
											for n4 = 9; n4 > 0; n4-- {
												for n3 = 9; n3 > 0; n3-- {
													for n2 = 9; n2 > 0; n2-- {
														for n1 = 9; n1 > 0; n1-- {

															inqueue = []int64{n14, n13, n12, n11, n10, n9, n8, n7, n6, n5, n4, n3, n2, n1}

															// fmt.Print(inqueue)
															for _, l := range lines {
																cmd := strings.Split(l, " ")

																switch cmd[0] {
																case "inp":
																	inp(cmd[1], &inqueue)
																case "add":
																	add(cmd[1], cmd[2])
																case "mul":
																	mul(cmd[1], cmd[2])
																case "div":
																	div(cmd[1], cmd[2])
																case "mod":
																	mod(cmd[1], cmd[2])
																case "eql":
																	eql(cmd[1], cmd[2])
																}
															}
															if z == 0 {
																fmt.Printf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",
																	n14, n13, n12, n11, n10, n9, n8, n7, n6, n5, n4, n4, n2, n1)
																fmt.Printf(" Final: w=%f, x=%f, y=%f, z=%f\n", w, x, y, z)
																fmt.Printf("DONE!\n")
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

}
