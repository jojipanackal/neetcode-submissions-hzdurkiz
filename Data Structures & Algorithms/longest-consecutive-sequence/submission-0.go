func longestConsecutive(nums []int) int {
    var result int
    in_nums := make(map[int]bool)

    for _, num := range nums {
        in_nums[num] = true
    }

    n := len(nums)
    for _, num := range nums {
        count := 0
        if !in_nums[num-1] {
            count = 1;
            for i:=1; i<n; i++ {
                if in_nums[num+i] {
                    count++
                } else {
                    break
                }
            }
            if count > result {
                result = count
            }
        }
    }
    return result
}
