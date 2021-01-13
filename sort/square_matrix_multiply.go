package sort

//蛮力法
func squareMatrixMultiply1(A, B [][]int) (C [][]int) {
	C = make([][]int, len(A))
	for i, _ := range C {
		C[i] = make([]int, len(B[0]))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[i]); j++ {
			for k := 0; k < len(B); k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

//分治法
func squareMatrixMultiply2(A, B [][]int, n int) (C [][]int) {
	//初始化矩阵C
	C = make([][]int, n)
	for i, _ := range C {
		C[i] = make([]int, n)
	}
	if n == 1 {
		C=squareMatrixMultiply1(A,B)
		return C
	}
	if n > 1 {
		m := n / 2
		A1 := divideMatrix(A, 1, 1)
		A2 := divideMatrix(A, 1, 2)
		A3 := divideMatrix(A, 2, 1)
		A4 := divideMatrix(A, 2, 2)
		B1 := divideMatrix(B, 1, 1)
		B2 := divideMatrix(B, 1, 2)
		B3 := divideMatrix(B, 2, 1)
		B4 := divideMatrix(B, 2, 2)
		C1 := addTwoMatrix(squareMatrixMultiply2(A1, B1, m), squareMatrixMultiply2(A2, B3, m), m)
		C2 := addTwoMatrix(squareMatrixMultiply2(A1, B2, m), squareMatrixMultiply2(A2, B4, m), m)
		C3 := addTwoMatrix(squareMatrixMultiply2(A3, B1, m), squareMatrixMultiply2(A4, B3, m), m)
		C4 := addTwoMatrix(squareMatrixMultiply2(A3, B2, m), squareMatrixMultiply2(A4, B4, m), m)

		C = togetherMatrix(C1, C2, C3, C4, m)
	}
	return C
}

//截取A11 A12	四个子矩阵
//	  A21 A22
func divideMatrix(A [][]int, x, y int) (C [][]int) {
	a := len(A) / 2
	if x == 1 && y == 1 {
		for i := 0; i < a; i++ {
			C = append(C, A[i])
			C[i] = C[i][:a]
		}
	}
	if x == 1 && y == 2 {
		for i := 0; i < a; i++ {

			C = append(C, A[i])
			C[i] = C[i][a:]
		}
	}
	if x == 2 && y == 1 {
		for i := a; i < len(A); i++ {
			C = append(C, A[i])
			C[i-a] = C[i-a][:a]
		}
	}
	if x == 2 && y == 2 {
		for i := a; i < len(A); i++ {
			C = append(C, A[i])
			C[i-a] = C[i-a][a:]
		}
	}
	return C
}

//四个子矩阵拼接成一个矩阵
func togetherMatrix(a, b, c, d [][]int, n int) (C [][]int) {
	C = make([][]int, 2*n)
	for i, _ := range C {
		C[i] = make([]int, 2*n)
	}
	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			if i < n {
				if j < n {
					C[i][j] = a[i][j]
				} else {
					C[i][j] = b[i][j-n]
				}
			} else {
				if j < n {
					C[i][j] = c[i-n][j]
				} else {
					C[i][j] = d[i-n][j-n]
				}
			}

		}
	}
	return C
}

//两个矩阵相加
func addTwoMatrix(A, B [][]int, n int) (C [][]int) {
	C = make([][]int, n)
	for i, _ := range C {
		C[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}