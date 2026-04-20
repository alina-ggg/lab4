package main

import (
	"fmt"
	"math"
	"strings"
)

type iterationStep struct {
	N    int
	A    float64
	B    float64
	X    float64
	Diff float64
}

func f(x float64) float64 {
	return math.Cos(x+0.3) - math.Pow(x, 2)
}

func df(x float64) float64 {
	return -math.Sin(x+0.3) - 2*x
}

func phi(x float64) float64 {
	return math.Sqrt(math.Cos(x + 0.3))
}

func printBorder(width int) {
	fmt.Printf("+%s\n", strings.Repeat("-", width))
}

func printTable(title string, history []iterationStep) {
	tableWidth := 55
	fmt.Printf("\n=== %s ===\n", title)

	printBorder(tableWidth)

	// Печать заголовка
	fmt.Printf("| %-6s%-16s%-16s%-17s|\n", "N", "a_n (x_n)", "b_n (x_n+1)", "diff")

	// Разделитель шапки
	fmt.Printf("| %s  |\n", strings.Repeat("-", tableWidth-2))

	for _, step := range history {
		fmt.Printf("| %-6d%-16.6f%-16.6f%-17.6f|\n",
			step.N, step.A, step.B, step.Diff)
	}

	printBorder(tableWidth)
}

func main() {
	a, b := 0.0, 1.0
	eps := 1e-4

	// Метод половинного деления
	var bisection []iterationStep
	currA, currB := a, b
	k := 0
	for (currB - currA) > eps {
		mid := (currA + currB) / 2.0
		k++
		bisection = append(bisection, iterationStep{k, currA, currB, mid, currB - currA})
		if f(currA)*f(mid) < 0 {
			currB = mid
		} else {
			currA = mid
		}
	}
	printTable("Метод половинного деления", bisection)

	// Метод Ньютона
	var newton []iterationStep
	xPrev := 0.5
	k = 0
	for {
		xNext := xPrev - f(xPrev)/df(xPrev)
		diff := math.Abs(xNext - xPrev)
		k++
		newton = append(newton, iterationStep{k, xPrev, xNext, xNext, diff})
		if diff < eps {
			break
		}
		xPrev = xNext
	}
	printTable("Метод Ньютона", newton)

	// Метод простых итераций
	var fixedPoint []iterationStep
	xPrev = 0.5
	k = 0
	for {
		xNext := phi(xPrev)
		diff := math.Abs(xNext - xPrev)
		k++
		fixedPoint = append(fixedPoint, iterationStep{k, xPrev, xNext, xNext, diff})
		if diff < eps {
			break
		}
		xPrev = xNext
	}
	printTable("Метод простых итераций", fixedPoint)
}
