package main

import (
	"fmt"
	"flag"
	"strconv"
)

func init_vector (vector *[][]int, fils int, cols int) {
	for i := 0; i < fils; i++ {
		for j := 0; j < cols; j++ {
			(*vector)[i][j] = 1;
		}
	}
}

func print_vector (vector *[][]int, fils int, cols int) {
	for i := 0; i < fils; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", (*vector)[i][j]);
		}
	}
}

func main () {

	flag.Parse()

	if (flag.NArg() > 2) {
		fmt.Printf("\nsequential_vector %s %s %s\n", flag.Arg(0), flag.Arg(1), flag.Arg(2))

		vector1_fils,_ := strconv.Atoi(flag.Arg(0))
		vector1_cols,_ := strconv.Atoi(flag.Arg(1))
		vector2_fils := vector1_cols;
		vector2_cols,_ := strconv.Atoi(flag.Arg(2))

		vector1 := make([][]int, vector1_fils)
		for i := 0; i < vector1_fils; i++ {
			vector1[i] = make([]int, vector1_cols)
		}

		vector2 := make([][]int, vector2_fils)
		for i := 0; i < vector2_fils; i++ {
			vector2[i] = make([]int, vector2_cols)
		}

		vectorR := make([][]int, vector1_fils)
		for i := 0; i < vector1_fils; i++ {
			vectorR[i] = make([]int, vector2_cols)
		}

		init_vector(&vector1, vector1_fils, vector1_cols)
		init_vector(&vector2, vector2_fils, vector2_cols)

		// Bucle principal
		var acum int

		for i := 0; i < vector1_fils; i++ {
			for j := 0; j < vector2_cols; j++ {
				acum = 0
				for k := 0; k < vector1_cols; k++ {
					acum += vector1[i][k] * vector2[k][j]
				}
				vectorR[i][j] = acum
			}
		}

		//print_vector(&vectorR, vector1_fils, vector2_cols)
	}
}

