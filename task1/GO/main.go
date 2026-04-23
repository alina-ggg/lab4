package main

import (
	"fmt"
	"math"
	"strings"
)

// структура для хранения шага итерации
type IterationStep struct {
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
	// Для второго корня : -math.Sqrt(...)
	return math.Sqrt(math.Cos(x + 0.3))
}

// Метод половинного деления
func solveBisection(a, b, eps float64) []IterationStep {
	var history []IterationStep
	currA, currB := a, b
	k := 0
	for math.Abs(currB-currA) > eps {
		mid := (currA + currB) / 2.0
		k++
		history = append(history, IterationStep{
			N:    k,
			A:    currA,
			B:    currB,
			X:    mid,
			Diff: math.Abs(currB - currA),
		})
		if f(currA)*f(mid) < 0 {
			currB = mid
		} else {
			currA = mid
		}
	}
	return history
}

// Метод Ньютона
func solveNewton(x0, eps float64) []IterationStep {
	var history []IterationStep
	xPrev := x0
	k := 0
	for {
		xNext := xPrev - f(xPrev)/df(xPrev)
		diff := math.Abs(xNext - xPrev)
		k++
		history = append(history, IterationStep{
			N:    k,
			A:    xPrev,
			B:    xNext,
			X:    xNext,
			Diff: diff,
		})
		if diff < eps {
			break
		}
		xPrev = xNext
	}
	return history
}

// Метод простых итераций
func solveFixedPoint(x0, eps float64) []IterationStep {
	var history []IterationStep
	xPrev := x0
	k := 0
	for {
		xNext := phi(xPrev)
		diff := math.Abs(xNext - xPrev)
		k++
		history = append(history, IterationStep{
			N:    k,
			A:    xPrev,
			B:    xNext,
			X:    xNext,
			Diff: diff,
		})
		if diff < eps {
			break
		}
		xPrev = xNext
	}
	return history
}

func printBorder(width int) {
	fmt.Printf("+%s+\n", strings.Repeat("-", width))
}

func printTable(title string, history []IterationStep) {
	const tableWidth = 55
	fmt.Printf("\n=== %s ===\n", title)
	printBorder(tableWidth)

	// Заголовок
	fmt.Printf("| %-6s%-16s%-16s%-17s|\n", "N", "a_n (x_n)", "b_n (x_n+1)", "diff")

	// Разделитель шапки
	fmt.Printf("| %s |\n", strings.Repeat("-", tableWidth-2))

	for _, step := range history {
		fmt.Printf("| %-6d%-16.6f%-16.6f%-17.6f|\n",
			step.N, step.A, step.B, step.Diff)
	}
	printBorder(tableWidth)
}

func main() {
	// Параметры для первого корня
	a, b := 0.0, 1.0
	x0 := 0.5
	eps := 1e-4

	/* Для второго корня:
	a, b = -2.0, 0.0
	x0 = -0.8
	минус в функцию phi()
	*/

	printTable("Метод половинного деления", solveBisection(a, b, eps))
	printTable("Метод Ньютона", solveNewton(x0, eps))
	printTable("Метод простых итераций", solveFixedPoint(x0, eps))
}
