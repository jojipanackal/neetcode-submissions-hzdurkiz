func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

    l, r := 0, m*n-1

    for l <=r {
        mid := l + (r-l)/2

        row := mid/n
        col := mid%n 

        if target > matrix[row][col] {
            l = mid+1
        } else if target < matrix[row][col] {
            r = mid-1
        } else {
            return true
        }
    }

    return false
}
