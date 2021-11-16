package main

import (
	"bufio"
	"fmt"
	"os"

	vectors "github.com/husseinelguindi/physics11-vectors"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	simpleVecs := []vectors.SimpleVector{}
	complexVecs := []vectors.Vector{}
	totalVecs := 0

Loop:
	for {
		fmt.Println("1. Add simple vector (ex. 20.1 [W])")
		fmt.Println("2. Add complex vector (ex. 20.1 [N 19.2 E])")
		if totalVecs >= 2 {
			fmt.Printf("3. Calculate inputted vectors (%d)", totalVecs)
		}
		choice := readInt(scanner, 10, 32, "\nYour choice: ")

		switch choice {
		case 1:
			simpleVecs = append(simpleVecs, getSimpleVec(scanner))
			totalVecs++
		case 2:
			complexVecs = append(complexVecs, getComplexVec(scanner))
			totalVecs++
		case 3:
			if totalVecs >= 2 {
				break Loop
			}
			fallthrough
		default:
			fmt.Print("Invalid choice.\n\n")
			continue
		}
		fmt.Print("\n\n")
	}

	for _, cv := range complexVecs {
		r1, r2 := cv.Resolve()
		simpleVecs = append(simpleVecs, r1, r2)
	}
	displayAddVec(false, simpleVecs...)

InvertLoop:
	for {
		choice := readRune(scanner, "\nInvert resultant vector angle? (y/n) ")
		switch choice {
		case 'y', 'Y':
			displayAddVec(true, simpleVecs...)
			break InvertLoop
		case 'n', 'N':
			break InvertLoop
		default:
			fmt.Print("Invalid choice.\n\n")
		}
	}
}

func displayAddVec(inverseAngle bool, simpleVectors ...vectors.SimpleVector) {
	rv := vectors.Add(inverseAngle, simpleVectors...)
	if sv, ok := rv.Simplify(); ok {
		fmt.Println(sv)
	} else {
		fmt.Println(rv)
	}
}
