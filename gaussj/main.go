package main

import (
	"fmt"
	"log"
	"math"
)

func partial_pivoting(a [][]float64, b []float64, n_row, n_col int) {
	/* forward elimination */
	for i := 0; i < n_col; i++ {
		max := 0.0
		var pivot_j int

		/* this routine is just partial */
		for j := i; j < n_row; j++ {
			if max < math.Abs(a[j][i]) {
				max = math.Abs(a[j][i])
				pivot_j = j
			}
		}

		if max <= 0.0000000001 {
			fmt.Println(a)
			fmt.Println(b)
			log.Fatalln("maybe singler matrix")
		}

		if i != pivot_j {
			for k := i; k < n_col; k++ {
				a[i][k], a[pivot_j][k] = a[pivot_j][k], a[i][k]
			}
			b[i], b[pivot_j] = b[pivot_j], b[i]
		}

		pivinv := 1.0 / a[i][i]
		a[i][i] = 1.0
		for k := i + 1; k < n_col; k++ {
			a[i][k] *= pivinv
		}
		b[i] *= pivinv

		for j := i + 1; j < n_row; j++ {
			for k := i + 1; k < n_col; k++ {
				a[j][k] -= a[j][i] * a[i][k]
			}
			b[j] -= b[i] * a[j][i]
			a[j][i] = 0
		}
	}

	/* backward substitution */
	var ans []float64
	for i := 0; i < n_row; i++ {
		row_i := n_row - 1 - i
		val := b[row_i]

		for j := 0; j < i; j++ {
			col_i := n_col - 1 - j
			val -= a[row_i][col_i] * ans[j]
		}
		ans = append(ans, val)
	}

	fmt.Print("ans: [")
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v", ans[len(ans)-1-i])
		if i == len(ans)-1 {
			fmt.Print("]")
		} else {
			fmt.Print(", ")
		}
	}
}

func main() {
	var n_row, n_col int

	/* want to input by passing filepath */
	fmt.Print("input #row: ")
	fmt.Scan(&n_row)
	fmt.Print("input #col: ")
	fmt.Scan(&n_col)

	if n_col <= 0 || n_row <= 0 {
		log.Fatalln("both #row and #col should be more than 0.")
	}

	a := make([][]float64, n_row)
	for i := 0; i < n_row; i++ {
		a[i] = make([]float64, n_col)
		for j := 0; j < n_col; j++ {
			fmt.Printf("input A[%d][%d]: ", i, j)
			fmt.Scan(&a[i][j])
		}
	}

	b := make([]float64, n_col)
	for i := 0; i < n_col; i++ {
		fmt.Printf("input b[%d]: ", i)
		fmt.Scan(&b[i])
	}

	partial_pivoting(a, b, n_row, n_row)
}
