package main

import (
	"bufio"
	"fmt"
	"strconv"

	vectors "github.com/husseinelguindi/physics11-vectors"
)

func readFloat(scanner *bufio.Scanner, bitsize int, msg string) float64 {
	fmt.Print(msg)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), bitsize)
		if err == nil {
			return f
		}
		fmt.Print(msg)
	}
	return 0
}

func readInt(scanner *bufio.Scanner, base, bitsize int, msg string) int64 {
	fmt.Print(msg)
	for scanner.Scan() {
		f, err := strconv.ParseInt(scanner.Text(), base, bitsize)
		if err == nil {
			return f
		}
		fmt.Print(msg)
	}
	return 0
}

func readRune(scanner *bufio.Scanner, msg string) rune {
	fmt.Print(msg)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 1 {
			return rune(text[0])
		}
		fmt.Print(msg)
	}
	return 0
}

func readDirection(scanner *bufio.Scanner, msg string) vectors.Direction {
	for {
		d := vectors.Direction(readRune(scanner, msg))
		if d.Valid() {
			return d
		}
		fmt.Print(msg)
	}
}

func getComplexVec(scanner *bufio.Scanner) vectors.Vector {
	fmt.Println("\nInputting a vector:")
	return vectors.Vector{
		Mag:             readFloat(scanner, 64, "Magnitude: "),
		StartDirection:  readDirection(scanner, "Starting direction (N/E/S/W): "),
		RelativeAngle:   readFloat(scanner, 64, "Relative angle: "),
		TowardDirection: readDirection(scanner, "Towards direction (N/E/S/W): "),
	}
}

func getSimpleVec(scanner *bufio.Scanner) vectors.SimpleVector {
	return vectors.SimpleVector{
		Mag:       readFloat(scanner, 64, "\nMagnitude: "),
		Direction: readDirection(scanner, "Direction (N/E/S/W): "),
	}
}
