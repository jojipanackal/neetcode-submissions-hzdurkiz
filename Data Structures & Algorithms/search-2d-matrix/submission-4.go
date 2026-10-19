func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) // rows
	n := len(matrix[0]) // columns
	midr := m/2
	l := 0
	r := n-1
	
	for l <= r {
		if target == matrix[midr][l] || target == matrix[midr][r] {
			return true
		}

		if target < matrix[midr][0] {
			if midr == 0 {
				return false
			}
			if target > matrix[midr-1][n-1] {
				return false
			}
			midr--
		} else if target > matrix[midr][n-1] {
			if midr == m-1 {
				return false
			}
			if target < matrix[midr+1][0] {
				return false
			}
			midr++
		} else {
			mid := l + (r-l)/2

			if target > matrix[midr][mid] {
				l = mid+1
			} else if target < matrix[midr][mid] {
				r = mid-1
			} else {
				return true
			}
		}
	}

	return false
}
